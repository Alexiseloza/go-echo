package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-echo/database"
	"go-echo/models"
)

// Create Category
func Category_post(c echo.Context) error {
	var body models.CategoryModel
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		r := map[string]string{
			"status":   "Error",
			"Response": "Response Error",
			"Message":  "User Created Can't be Created",
			"User":     "The data isn't correct",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set("App-Name", "NameHere")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(r)
	}
	if len(body.Name) == 0 {
		r := map[string]string{
			"status":    "Error",
			"Response":  "Field Missing",
			"Message":   "You must fill all the fields",
			"fieldname": "",
		}
		if len(body.Name) == 0 {
			r["fieldname"] = "Category Name"
		}
		c.Response().WriteHeader(http.StatusUnprocessableEntity) // https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/42
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set("App-Name", "NameHere")
	}
	//save DB
	register := bson.D{{"Name", body.Name}, {"slug", slug.Make(body.Name)}}
	database.CategoryColection.InsertOne(context.TODO(), register)
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"Message":  "Category Created Successfully",
		"Headers":  c.Request().Header.Get("Authorization"),
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(r)

}

// all Categories
func GetAllCategories(c echo.Context) error {

	coursor, err := database.CategoryColection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", +1}})) // set -1 for descend mode
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = coursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return c.JSON(http.StatusOK, results)
}

// Category BY ID
func GetCategoryById(c echo.Context) error {
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var results bson.M
	if err := database.CategoryColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&results); err != nil {
		r := map[string]string{
			"status":    "Error",
			"Response":  "Category NOT Found",
			"Message":   "Please Check the category ID",
			"fieldname": "",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set("App-Name", "NameHere")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(r)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return c.JSON(http.StatusOK, results)

}

// Update category
func UpdateCategory(c echo.Context) error {
	var body models.CategoryModel
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		response := map[string]string{
			"status":   "Error",
			"Messagge": "Invalid DATA inthis Category",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(body.Name) == 0 {
		response := map[string]string{
			"status":   "Error",
			"Messagge": "Invalid ID or Category",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return c.JSON(http.StatusBadRequest, response)
	}
	var result bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.CategoryColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result); err != nil {
		response := map[string]string{
			"status":   "Error",
			"Messagge": "Invalid DATA or  Category",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return c.JSON(http.StatusInternalServerError, response)

	}
	//Edition
	update := make(map[string]interface{})
	update["Name"] = body.Name
	update["slug"] = slug.Make(body.Name)
	updateString := bson.M{
		"$set": update,
	}
	database.CategoryColection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}}, updateString)
	result = update
	response := map[string]interface{}{
		"status":  "Success",
		"Data":    result,
		"Message": "The category has been successfully updated.",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(http.StatusOK, response)

}

// Delete category
func DeleteCategory(c echo.Context) error {
	var result bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.CategoryColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result); err != nil {
		response := map[string]string{
			"status":   "Error",
			"Messagge": "Invalid ID Please Check!!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
		return c.JSON(http.StatusInternalServerError, response)

	}
	database.CategoryColection.DeleteOne(context.TODO(), bson.M{"_id": objID})

	response := map[string]interface{}{
		"status":  "Success",
		"Data":    result,
		"Message": "The category has been DELETED successfully.",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSON(http.StatusOK, response)

}

//end
