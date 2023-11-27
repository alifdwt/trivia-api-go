package main

import (
	"fmt"
	"trivia-api-go/database"
	"trivia-api-go/pkg/postgres"
	"trivia-api-go/routes"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()

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