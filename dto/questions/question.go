package questionsdto

type CreateQuestionRequest struct {
	Question string `json:"question" form:"question" validate:"required" example:"What is your name?"`
	Answer string `json:"answer" form:"answer" validate:"required" example:"John Doe"`
	WrongAnswer1 string `json:"wrong_answer_1" form:"wrong_answer_1" validate:"required" example:"Jane Doe"`
	WrongAnswer2 string `json:"wrong_answer_2" form:"wrong_answer_2" validate:"required" example:"Jack Doe"`
	WrongAnswer3 string `json:"wrong_answer_3" form:"wrong_answer_3" validate:"required" example:"Jill Doe"`
}

type UpdateQuestionRequest struct {
	Question string `json:"question" form:"question" validate:"required" example:"What is your name?"`
	Answer string `json:"answer" form:"answer" validate:"required" example:"John Doe"`
	WrongAnswer1 string `json:"wrong_answer_1" form:"wrong_answer_1" validate:"required" example:"Jane Doe"`
	WrongAnswer2 string `json:"wrong_answer_2" form:"wrong_answer_2" validate:"required" example:"Jack Doe"`
	WrongAnswer3 string `json:"wrong_answer_3" form:"wrong_answer_3" validate:"required" example:"Jill Doe"`
}

type QuestionResponse struct {
	ID int `json:"id" example:"1"`
	Question string `json:"question" example:"What is your name?"`
	Answer string `json:"answer" example:"John Doe"`
	WrongAnswer1 string `json:"wrong_answer_1" example:"Jane Doe"`
	WrongAnswer2 string `json:"wrong_answer_2" example:"Jack Doe"`
	WrongAnswer3 string `json:"wrong_answer_3" example:"Jill Doe"`
}