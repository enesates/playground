package db

import (
	"errors"
	"log"
	"time"

	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func CheckUserExists(username string, email string) bool {
	user := models.User{
		Username: username,
		Email:    email,
	}

	result := GormDB.Where("username = ? OR email = ?", user.Username, user.Email).First(&user)
	return result.RowsAffected > 0
}

func GetOrCreateSession(userID string) (models.Session, error) {
	session := models.Session{}
	GormDB.Where("user_id = ?", userID).First(&session)

	isSessionExpired := session.ID != "" && session.ExpiresAt.Before(time.Now())

	if session.ID == "" || isSessionExpired {
		if isSessionExpired {
			result := GormDB.Delete(&session)

			if result.Error != nil {
				log.Printf("Error deleting session: %v", result.Error)
				return session, errors.New("Internal Server Error")
			}
		}

		expDate := time.Now().Add(time.Hour * 24)
		token, err := helpers.CreateToken(userID, expDate)

		if err != nil {
			log.Printf("Error logging user: %v", err)
			return session, errors.New("Internal Server Error")
		}

		session = models.Session{
			ID:        helpers.GetUUID(),
			UserID:    userID,
			Token:     token,
			ExpiresAt: expDate,
		}

		result := GormDB.Create(&session)
		if result.Error != nil {
			log.Printf("Error creating session: %v", result.Error)
			return session, errors.New("Internal Server Error")
		}
	}

	return session, nil

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

	result := GormDB.Create(&user)

	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return nil, errors.New("Internal Server Error")
	}

	log.Printf("User created successfully with ID: %s", user.ID)
	return &user, nil
}

func GetUser(userDTO models.UserLoginDTO) (models.User, models.Session, error) {
	var session = models.Session{}

	user := models.User{
		Email: userDTO.Email,
	}

	result := GormDB.Where("Email = ?", user.Email).First(&user)

	if result.Error != nil {
		log.Printf("Error logging user: %v", result.Error)
		return user, session, errors.New("Internal Server Error")
	} else if !helpers.ComparePassword(user.PasswordHash, userDTO.Password) {
		log.Printf("Credentials are not correct")
		return user, session, errors.New("Invalid credentials")
	}

	session, err := GetOrCreateSession(user.ID)

	return user, session, err
}
