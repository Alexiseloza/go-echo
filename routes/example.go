package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"go-echo/models"
)

/* Release 1.0 */

func Test_get(c echo.Context) error {
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"Headers":  c.Request().Header.Get("Authorization"),
	}
	// c.Response().Header().Set("MyHeader", "yourSite")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(200) // HTTP Status Code: 200 - OK
	c.Response().Header().Set("App-Name", "NameHere")

	return json.NewEncoder(c.Response()).Encode(r)
}
func Test_get_userID(c echo.Context) error {
	id := c.Param("userId")
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"ID":       "ID=" + id,
		"Headers":  c.Request().Header.Get("Authorization"),
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	return json.NewEncoder(c.Response()).Encode(r)

}
func Test_post(c echo.Context) error {
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
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"Message":  "User Created Successfully",
		"User":     body.Name,
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(r)

}

//	func Test_post(c echo.Context) error {
//		r := map[string]string{
//			"status":   "success",
//			"Response": "Response OK",
//			"Message":  "User Created Successfully",
//		}
//		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
//		return json.NewEncoder(c.Response()).Encode(r)
//	}
func Test_put(c echo.Context) error {
	id := c.Param("userId")
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"ID":       "ID=" + id,
		"Message":  "Data updated successfully!",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return json.NewEncoder(c.Response()).Encode(r)
}
func Test_delete(c echo.Context) error {
	id := c.Param("userId")
	r := map[string]string{
		"status":   "success",
		"Response": "Response OK",
		"ID":       "ID=" + id,
		"Message":  "User deleted successfully!",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return json.NewEncoder(c.Response()).Encode(r)
}

// query string
func Search_get(c echo.Context) error {
	id := c.QueryParam("id")
	slug := c.QueryParam("slug") // make a URL-friendly slug from the name parameter
	r := map[string]string{
		"status":    "success",
		"Response":  "result OK",
		"result ID": "ID=" + id,
		"slug":      "SLUG=" + slug,
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return json.NewEncoder(c.Response()).Encode(r)
}

func Upload_file(c echo.Context) error {
	file, err := c.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	src, _ := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// rename file
	var extension = strings.Split(file.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	photo := string(time[4][6:14]) + "." + extension
	var fileupload = "./public/uploads/" + photo
	// create file
	dist, err := os.Create(fileupload)
	if err != nil {
		return err
	}
	defer dist.Close()
	if _, err = io.Copy(dist, src); err != nil {
		return err
	}
	r := map[string]string{
		"status":   "success",
		"Response": "Upload Successfully",
		"Image":    "Image=" + photo,
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("App-Name", "NameHere")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(r)

}
