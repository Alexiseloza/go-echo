package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-echo/database"
	"go-echo/models"
)

// Get ALL
func GetAllProducts(c echo.Context) error {
	pipeline := []bson.M{
		bson.M{"$match": bson.M{}},
		bson.M{"$lookup": bson.M{"from": "categories", "localField": "category_id", "foreignField": "_id", "as": "categories"}},
		bson.M{"$sort": bson.M{"_id": -1}},
	}
	cursor, err := database.ProductColection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unsespected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unsespected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

// Get product by ID
func GetProductById(c echo.Context) error {
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": objID}},
		bson.M{"$lookup": bson.M{"from": "categories", "localField": "category_id", "foreignField": "_id", "as": "categories"}},
		bson.M{"$sort": bson.M{"_id": -1}},
	}
	cursor, err := database.ProductColection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unsespected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unsespected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result[0])

}

// add product
func AddProduct(c echo.Context) error {
	var body models.ProductModel
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unsespected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(body.Name) == 0 {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "The Name is Required!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	//save on the DB
	CatID, _ := primitive.ObjectIDFromHex(body.CategoryId)
	register := bson.D{{"name", body.Name}, {"slug", slug.Make(body.Name)}, {"precio", body.Price}, {"stock", body.Stock}, {"descripcion", body.Description}, {"categoria_id", CatID}}
	database.ProductColection.InsertOne(context.TODO(), register)

	//retorno respuesta
	response := map[string]string{
		"Status":   "OK",
		"Messagge": "Product saved successfully",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(response)
}

// update product
func UpdateProduct(c echo.Context) error {
	var body models.ProductModel
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Unexpected Error!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(body.Name) == 0 {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "The Name is Empty!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)

	}
	var result bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param(("id")))
	if err := database.ProductColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "Error with the ID",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)
	}
	//edit product
	category_id, _ := primitive.ObjectIDFromHex(body.CategoryId)
	register := make(map[string]interface{})
	register["name"] = body.Name
	register["price"] = body.Price
	register["description"] = body.Description
	register["slug"] = slug.Make(body.Name)
	register["isPromo"] = body.CategoryId
	register["category_id"] = category_id
	updateString := bson.M{
		"$set": register,
	}
	database.ProductColection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}}, updateString)
	response := map[string]string{
		"Status":  "OK",
		"Message": "Product Updated Successfully",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)
}

// delete product
func DeleteProduct(c echo.Context) error {
	var result bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	if err := database.ProductColection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result); err != nil {
		response := map[string]string{
			"Status":   "Error",
			"Messagge": "The id is not valid.",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, response)

	}
	database.ProductColection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	response := map[string]string{
		"Status":  "OK",
		"Message": "Product Deleted Successfully"}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)
}

// End
