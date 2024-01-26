package handlers

import (
	"net/http"
	"strconv"
	"time"
	questionsdto "trivia-api-go/dto/questions"
	dto "trivia-api-go/dto/results"
	"trivia-api-go/models"
	"trivia-api-go/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerQuestion struct {
	QuestionRepository repositories.QuestionRepository
}

func NewQuestionHandler(QuestionRepository repositories.QuestionRepository) *handlerQuestion {
	return &handlerQuestion{QuestionRepository}
}

// @Tags Question
// @Summary Get All Questions
// @Router /question [get]
// @Produce json
// @Success 200 {object} dto.SuccessResult{data=[]models.Question}
// @Failure 500 {object} dto.ErrorResult
func (h *handlerQuestion) FindAll(c echo.Context) error {
	questions, err := h.QuestionRepository.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: questions})
}

// @Tags Question
// @Summary Get Question by ID
// @Router /question/{id} [get]
// @Produce json
// @Success 200 {object} dto.SuccessResult{data=models.Question}
// @Failure 500 {object} dto.ErrorResult
func (h *handlerQuestion) FindByID(c echo.Context) error {
	questionId, _ := strconv.Atoi(c.Param("questionId"))

	question, err := h.QuestionRepository.FindByID(questionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: question})
}

// @Tags Question
// @Summary Create Question
// @Router /question [post]
// @Param question body questionsdto.CreateQuestionRequest true "Create Question"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult{data=models.Question}
// @Failure 500 {object} dto.ErrorResult
func (h *handlerQuestion) Create(c echo.Context) error {
	request := new(questionsdto.CreateQuestionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	question := models.Question{
		Question: request.Question,
		Answer: request.Answer,
		WrongAnswer1: request.WrongAnswer1,
		WrongAnswer2: request.WrongAnswer2,
		WrongAnswer3: request.WrongAnswer3,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	data, err := h.QuestionRepository.Create(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

// @Tags Question
// @Summary Update Question
// @Router /question/{id} [put]
// @Param id path int true "Question ID"
// @Param question body questionsdto.UpdateQuestionRequest true "Update Question"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult{data=models.Question}
// @Failure 500 {object} dto.ErrorResult
func (h *handlerQuestion) Update(c echo.Context) error {
	request := new(questionsdto.UpdateQuestionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	questionId, _ := strconv.Atoi(c.Param("id"))
	question, err := h.QuestionRepository.FindByID(questionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Question != "" {
		question.Question = request.Question
	}

	if request.Answer != "" {
		question.Answer = request.Answer
	}

	if request.WrongAnswer1 != "" {
		question.WrongAnswer1 = request.WrongAnswer1
	}

	if request.WrongAnswer2 != "" {
		question.WrongAnswer2 = request.WrongAnswer2
	}

	if request.WrongAnswer3 != "" {
		question.WrongAnswer3 = request.WrongAnswer3
	}

	question.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	data, err := h.QuestionRepository.Update(question, questionId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

// @Tags Question
// @Summary Delete Question
// @Router /question/{id} [delete]
// @Param id path int true "Question ID"
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult{data=models.Question}
// @Failure 500 {object} dto.ErrorResult
func (h *handlerQuestion) Delete(c echo.Context) error {
	questionId, _ := strconv.Atoi(c.Param("questionId"))

	question, err := h.QuestionRepository.FindByID(questionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.QuestionRepository.Delete(question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}