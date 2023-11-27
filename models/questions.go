package models

type Question struct {
	ID int `json:"id" gorm:"primaryKey"`
	Question string `json:"question" gorm:"not null;type:varchar(255)"`
	Answer string `json:"answer" gorm:"not null;type:varchar(255)"`
	WrongAnswer1 string `json:"wrong_answer_1" gorm:"not null;type:varchar(255);column:wrong_answer_1"`
	WrongAnswer2 string `json:"wrong_answer_2" gorm:"not null;type:varchar(255);column:wrong_answer_2"`
	WrongAnswer3 string `json:"wrong_answer_3" gorm:"not null;type:varchar(255);column:wrong_answer_3"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (*Question) TableName() string {
	return "questions"
}