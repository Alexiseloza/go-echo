package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"go-echo/database"
	"go-echo/jwt"
	"go-echo/middleware"
	"go-echo/models"
	"go-echo/validator"
)

// Protected Router
func Protected(c echo.Context) error {
	if middleware.ValidateJWT(c) == 0 {
		response := map[string]string{
			"status":   "Error",
			"messagge": "Not AUTHORIZED",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	response := map[string]string{
		"status":   "OK",
		"messagge": "Route Protected Authorized!",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)

}

// Register
func UserRegister(c echo.Context) error {
	var body models.UserModel
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "Unespected Error",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if len(body.Name) == 0 {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Name is Required!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if len(body.Email) == 0 {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Email is Required!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if validator.Regex_email.FindStringSubmatch(body.Email) == nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Email is NOT Valid!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if validator.ValidatePassword(body.Password) == false {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Password must to have at least 6 characters and contain uppercase letters, lowercase letters, numbers and special characters.",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	var user bson.M
	if err := database.UsersCollection.FindOne(context.TODO(), bson.M{"email": body.Email}).Decode(&user); err == nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "An Error was ocurred",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	cost := 8
	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), cost)
	register := bson.D{{"name", body.Name}, {"phone", body.Phone}, {"password", string(bytes)}, {"email", body.Email}}
	database.UsersCollection.InsertOne(context.TODO(), register)

	response := map[string]string{
		"status":   "OK",
		"messagge": "The User has been created correctly.",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(response)

}

// User Login
func UserLogin(c echo.Context) error {
	var body models.LoginModel
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "An Error was ocurred",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if len(body.Email) == 0 {
		response := map[string]string{
			"status":   "Error",
			"messagge": "Email not provided",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if validator.Regex_email.FindStringSubmatch(body.Email) == nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Email is NOT correct",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	if validator.ValidatePassword(body.Password) == false {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Password must to have at least 6 characters and contain uppercase letters, lowercase letters, numbers and special characters.",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	var user bson.M
	if err := database.UsersCollection.FindOne(context.TODO(), bson.D{{"email", body.Email}}).Decode(&user); err != nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Email ins't correct , Please Check!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}
	passwordBytes := []byte(body.Password)
	passwordDB := []byte(user["password"].(string))
	errPassword := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if errPassword != nil {
		response := map[string]string{
			"status":   "Error",
			"messagge": "The Password ins't correct , Please Check!",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	} else {
		stringObjectID := user["_id"].(primitive.ObjectID).Hex()
		jwtKey, err := jwt.GenerateJWT(user["email"].(string), user["name"].(string), stringObjectID)
		if err != nil {
			response := map[string]string{
				"status":   "Error",
				"messagge": "An error ocurred  while generating the token: " + err.Error(),
			}
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			c.Response().WriteHeader(http.StatusBadRequest)
			return json.NewEncoder(c.Response()).Encode(response)
		} else {
			returned := models.TokenResponse{
				Name:  user["name"].(string),
				Token: jwtKey,
			}
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			c.Response().WriteHeader(http.StatusOK)
			return json.NewEncoder(c.Response()).Encode(returned)

		}
	}

}
