package database

import (
	"log"

	"ecommapi/internal/helpers.go"
)

func CreateUser(username string, email string, password string) {
	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		log.Println("User couldn't created. Error:", err)
		return
	}

	user := User{
		ID:           helpers.GetUUID(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	result := GormDB.Create(&user)

	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return
	}

	log.Printf("User created successfully with ID: %s", user.ID)
}
