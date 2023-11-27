package repositories

import (
	"trivia-api-go/models"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	FindAll() ([]models.Question, error)
	FindByID(id int) (models.Question, error)
	Create(question models.Question) (models.Question, error)
	Update(question models.Question, ID int) (models.Question, error)
	Delete(question models.Question) (models.Question, error)
}

type questionRepository struct{
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *questionRepository {
	return &questionRepository{db}
}

func (r *questionRepository) FindAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Find(&questions).Error

	return questions, err
}

func (r *questionRepository) FindByID(id int) (models.Question, error) {
	var question models.Question
	err := r.db.First(&question, id).Error

	return question, err
}

func (r *questionRepository) Create(question models.Question) (models.Question, error) {
	err := r.db.Create(&question).Error

	return question, err
}

func (r *questionRepository) Update(question models.Question, ID int) (models.Question, error) {
	err := r.db.Model(&question).Where("id = ?", ID).Updates(&question).Error

	return question, err
}

func (r *questionRepository) Delete(question models.Question) (models.Question, error) {
	err := r.db.Delete(&question).Error

	return question, err
}