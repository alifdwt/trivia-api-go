package database

import (
	"fmt"
	"trivia-api-go/models"
	"trivia-api-go/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.Question{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Success!")
}