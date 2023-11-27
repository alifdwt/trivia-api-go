package routes

import (
	"trivia-api-go/handlers"
	"trivia-api-go/pkg/postgres"
	"trivia-api-go/repositories"

	"github.com/labstack/echo/v4"
)

func QuestionRoutes(e *echo.Group) {
	questionRepository := repositories.NewQuestionRepository(postgres.DB)
	handler := handlers.NewQuestionHandler(questionRepository)

	e.GET("/question", handler.FindAll)
	e.GET("/question/:questionId", handler.FindByID)
	e.POST("/question", handler.Create)
	e.PUT("/question/:questionId", handler.Update)
	e.DELETE("/question/:questionId", handler.Delete)
}