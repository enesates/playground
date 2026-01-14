package db

import (
	"errors"

	db "ecommapi/internal/database"
	"ecommapi/internal/dtos"
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

func CreateUser(userDTO dtos.UserRegisterDTO) (*models.User, error) {
	hashedPassword, err := helpers.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:           helpers.GetUUID(),
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: string(hashedPassword),
		Role:         "customer",
	}

	if err := db.GormDB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserAndSession(userDTO dtos.UserLoginDTO) (*models.User, *models.Session, error) {
	user := models.User{
		Email: userDTO.Email,
	}

	if err := db.GormDB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return nil, nil, err
	} else if !helpers.ComparePassword(user.PasswordHash, userDTO.Password) {
		return nil, nil, errors.New("Invalid credentials")
	}

	session, err := GetOrCreateSession(user.ID)
	return &user, session, err
}

func GetUserByID(userID string) (*models.User, error) {
	user := models.User{}

	if err := db.GormDB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByToken(token string) (*models.User, error) {
	session, err := GetSessionByToken(token)
	if err != nil {
		return nil, err
	}

	return GetUserByID(session.UserID)
}
