package db

import (
	"errors"
	"log"

	db "ecommapi/internal/database"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func CheckUserExists(username string, email string) bool {
	user := models.User{
		Username: username,
		Email:    email,
	}

	result := db.GormDB.Where("username = ? OR email = ?", user.Username, user.Email).First(&user)
	return result.RowsAffected > 0
}

func CreateUser(userDTO models.UserRegisterDTO) (*models.User, error) {
	hashedPassword, err := helpers.HashPassword(userDTO.Password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, errors.New("Internal Server Error")
	}

	user := models.User{
		ID:           helpers.GetUUID(),
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: string(hashedPassword),
		Role:         "customer",
	}

	userResult := db.GormDB.Create(&user)

	if userResult.Error != nil {
		log.Printf("Error creating user: %v", userResult.Error)
		return nil, errors.New("Internal Server Error")
	}

	log.Printf("User created successfully with ID: %s", user.ID)
	return &user, nil
}

func GetUserAndSession(userDTO models.UserLoginDTO) (models.User, models.Session, error) {
	var session = models.Session{}

	user := models.User{
		Email: userDTO.Email,
	}

	userResult := db.GormDB.Where("email = ?", user.Email).First(&user)

	if userResult.Error != nil {
		log.Printf("Error logging user: %v", userResult.Error)
		return user, session, errors.New("Internal Server Error")
	} else if !helpers.ComparePassword(user.PasswordHash, userDTO.Password) {
		log.Printf("Credentials are not correct")
		return user, session, errors.New("Invalid credentials")
	}

	session, err := GetOrCreateSession(user.ID)

	return user, session, err
}

func GetUserByID(userID string) (models.User, error) {
	user := models.User{}
	userResult := db.GormDB.Where("id = ?", userID).First(&user)

	if userResult.Error != nil {
		log.Printf("Error getting user: %v", userResult.Error)
		return user, errors.New("Internal Server Error")
	}

	return user, nil
}

func GetUserByToken(token string) (models.User, error) {
	user := models.User{}

	session, err := GetSessionByToken(token)
	if err != nil {
		log.Printf("Error getting session: %v", err)
		return user, err
	}

	user, err = GetUserByID(session.UserID)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return user, err
	}

	return user, nil
}
