package questionsdto

type CreateQuestionRequest struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer string `json:"answer" form:"answer" validate:"required"`
	WrongAnswer1 string `json:"wrong_answer_1" form:"wrong_answer_1" validate:"required"`
	WrongAnswer2 string `json:"wrong_answer_2" form:"wrong_answer_2" validate:"required"`
	WrongAnswer3 string `json:"wrong_answer_3" form:"wrong_answer_3" validate:"required"`
}

type UpdateQuestionRequest struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer string `json:"answer" form:"answer" validate:"required"`
	WrongAnswer1 string `json:"wrong_answer_1" form:"wrong_answer_1" validate:"required"`
	WrongAnswer2 string `json:"wrong_answer_2" form:"wrong_answer_2" validate:"required"`
	WrongAnswer3 string `json:"wrong_answer_3" form:"wrong_answer_3" validate:"required"`
}

type QuestionResponse struct {
	ID int `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"`
	WrongAnswer1 string `json:"wrong_answer_1"`
	WrongAnswer2 string `json:"wrong_answer_2"`
	WrongAnswer3 string `json:"wrong_answer_3"`
}