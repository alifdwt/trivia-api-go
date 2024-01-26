package main

import (
	"fmt"
	"trivia-api-go/database"
	"trivia-api-go/pkg/postgres"
	"trivia-api-go/routes"

	_ "trivia-api-go/docs"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Trivia API
// @version         1.0
// @description     API for Trivia
// @termsOfService  http://swagger.io/terms/

// @contact.name   	API Support
// @contact.url    	http://www.swagger.io/support
// @contact.email  	FbV9T@example.com

// @license.name  	Apache 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @host      		localhost:5000
// @BasePath  		/api

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	postgres.DatabaseInit()
	database.RunMigration()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowedHeaders: []string{"*"},
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	routes.RouteInit(e.Group("/api"))

	fmt.Println("server running localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}