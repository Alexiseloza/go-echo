package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"go-echo/database"
	"go-echo/routes"
)

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1

var prefix string = "/api/v1/"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	// static files
	e.Static("/", "./public")
	// Database MOngo
	database.CheckConnection()
	e.GET(prefix+"docs/*any", echoSwagger.WrapHandler)

	// Router PATH

	e.GET(prefix+"poster", routes.Test_get) // route test

	// ROUTER Query Params  http:localhost:8080/api/v1/search?value1=2&value2=1
	e.GET(prefix+"search", routes.Search_get)
	//Authentication
	// authGroup := e.Group(prefix + "auth")
	// {
	// 	// authGroup.POST("register", routes.UserRegister) // Register User
	// 	// authGroup.POST("login", routes.UserLogin)       // Login User
	// 	// authGroup.POST("protected", routes.Protected)   // Log Out User

	// }
	e.POST(prefix+"auth/register", routes.UserRegister) // Register User
	e.POST(prefix+"auth/login", routes.UserLogin)       // Login User
	e.POST(prefix+"auth/protected", routes.Protected)   // Log Out User
	//User
	e.POST(prefix+"user", routes.Test_post)
	e.PUT(prefix+"user/:userId", routes.Test_put)
	e.DELETE(prefix+"user/:userId", routes.Test_delete)
	e.GET(prefix+"user/:id", routes.Test_get_userID)

	//Products
	e.GET(prefix+"products", routes.GetAllProducts)
	e.GET(prefix+"product/:id", routes.GetProductById)
	e.POST(prefix+"product", routes.AddProduct)
	e.PUT(prefix+"product/:id", routes.UpdateProduct)
	e.DELETE(prefix+"product/:id", routes.DeleteProduct)
	// product images
	e.GET(prefix+"product-image/:id", routes.GetImageProduct)
	e.POST(prefix+"product-images/:id", routes.UploadImageProduct)
	e.DELETE(prefix+"product-image/id", routes.DeletePhotoProduct)
	//Categories
	e.POST(prefix+"category", routes.Category_post)
	e.GET(prefix+"categories", routes.GetAllCategories)
	e.GET(prefix+"category/:id", routes.GetCategoryById)
	e.PUT(prefix+"category/:id", routes.UpdateCategory)
	e.DELETE(prefix+"category/remove/:id", routes.DeleteCategory)
	// Uploader
	e.POST(prefix+"upload", routes.Upload_file) // upload files to server

	// Environment
	erroVars := godotenv.Load()
	if erroVars != nil {
		panic(erroVars)

	}
	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
	// e.Logger.Fatal(e.Start(":8085"))
	fmt.Println(e.Server.Addr)

}
