package database

import (
	"errors"
	"log"

	"ecommapi/internal/helpers.go"
)

func CheckUserExists(username string, email string) bool {
	user := User{
		Username: username,
		Email:    email,
	}

	result := GormDB.First(&user)
	return result.RowsAffected > 1
}

func CreateUser(username string, email string, password string) error {
	if CheckUserExists(username, email) {
		log.Printf("User already exists")
		return errors.New("User already exists")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return errors.New("Internal Server Error")
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

		return errors.New("Internal Server Error")
	}

	log.Printf("User created successfully with ID: %s", user.ID)
	return nil
}

func GetUser(email string, password string) (User, error) {
	var err error

	user := User{
		Email: email,
	}

	result := GormDB.First(&user)

	if result.Error != nil {
		log.Printf("Error getting user: %v", result.Error)
		err = errors.New("Internal Server Error")
	} else if !helpers.ComparePassword(user.PasswordHash, password) {
		log.Printf("Credentials are not correct")
		err = errors.New("Invalid credentials")
	}

	return user, err
}
