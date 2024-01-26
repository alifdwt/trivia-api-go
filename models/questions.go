package models

type Question struct {
	ID int `json:"id" gorm:"primaryKey" example:"1"`
	Question string `json:"question" gorm:"not null;type:varchar(255)" example:"What is your name?"`
	Answer string `json:"answer" gorm:"not null;type:varchar(255)" example:"John Doe"`
	WrongAnswer1 string `json:"wrong_answer_1" gorm:"not null;type:varchar(255);column:wrong_answer_1" example:"Jane Doe"`
	WrongAnswer2 string `json:"wrong_answer_2" gorm:"not null;type:varchar(255);column:wrong_answer_2" example:"Jack Doe"`
	WrongAnswer3 string `json:"wrong_answer_3" gorm:"not null;type:varchar(255);column:wrong_answer_3" example:"Jill Doe"`
	CreatedAt string `json:"created_at" example:"2022-01-01 00:00:00"`
	UpdatedAt string `json:"updated_at" example:"2022-01-01 00:00:00"`
}

func (*Question) TableName() string {
	return "questions"
}