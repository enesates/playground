package db

import (
	"errors"
	"log"

	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func CheckUserExists(userDTO models.UserDTO) bool {
	user := models.User{
		Username: userDTO.Username,
		Email:    userDTO.Email,
	}

	result := GormDB.Where("Username = ? OR Email = ?", user.Username, user.Email).First(&user)
	return result.RowsAffected > 0
}

func CreateUser(userDTO models.UserDTO) error {
	hashedPassword, err := helpers.HashPassword(userDTO.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return errors.New("Internal Server Error")
	}

	user := models.User{
		ID:           helpers.GetUUID(),
		Username:     userDTO.Username,
		Email:        userDTO.Email,
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

func GetUser(userDTO models.UserDTO) (models.User, error) {
	var err error

	user := models.User{
		Email: userDTO.Email,
	}

	result := GormDB.Where("Email = ?", user.Email).First(&user)

	if result.Error != nil {
		log.Printf("Error getting user: %v", result.Error)
		err = errors.New("Internal Server Error")
	} else if !helpers.ComparePassword(user.PasswordHash, userDTO.Password) {
		log.Printf("Credentials are not correct")
		err = errors.New("Invalid credentials")
	}

	return user, err
}
