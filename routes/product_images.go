package routes

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-echo/database"
)

// Get product Image
func GetImageProduct(c echo.Context) error {
	var product bson.M
	objID, _ := primitive.ObjectIDFromHex("id")
	if err := database.ProductColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Error with the ID",
		}
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	coursor, err := database.ProductPhotoCollection.Find(context.TODO(), bson.D{{"product_id", objID}})
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = coursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return c.JSON(http.StatusOK, results)

}

// Upload Image  Product
func UploadImageProduct(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// rename file
	var extension = strings.Split(file.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	photo := string(time[4][6:14] + "," + extension)
	var thefile string = "public/uploads/products/" + photo

	dst, err := os.Create(thefile)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// save register in database
	var results bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.ProductColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&results); err != nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "Error image uploading.",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	register := bson.D{{"name", photo}, {"product_id", objID}}
	database.ProductPhotoCollection.InsertOne(context.TODO(), register)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	response := map[string]string{
		"status":   "Success",
		"messagge": "Image uploaded Successfully!",
	}
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(response)

}

func DeletePhotoProduct(c echo.Context) error {
	var result bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.ProductPhotoCollection.FindOne(context.TODO(), bson.M{
		"_id": objID,
	}).Decode(&result); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Something was wrong!!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	delete := "public/uploads/products/" + result["name"].(string)
	e := os.Remove(delete)
	if e != nil {
		log.Fatal(e)
	}
	database.ProductPhotoCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	response := map[string]string{
		"Status":   "OK",
		"Messagge": "The Image was deleted!!",
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)
}
