package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"go-career/component-go/swag/server"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/hello", server.Hello)
	e.GET("/goodbye", server.Hello)

	/*
		Or can use EchoWrapHandler func with configurations.
		url := echoSwagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
		e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
	*/
	e.Logger.Fatal(e.Start(":8080"))
}
