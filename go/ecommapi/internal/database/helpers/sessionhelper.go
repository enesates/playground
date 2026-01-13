package db

import (
	"errors"
	"log"
	"time"

	db "ecommapi/internal/database"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func GetSessionByToken(token string) (models.Session, error) {
	session := models.Session{
		Token: token,
	}

	sessionResult := db.GormDB.Where("token = ?", token).First(&session)

	if sessionResult.Error != nil {
		log.Printf("Error getting session: %v", sessionResult.Error)
		return session, errors.New("Internal Server Error")
	}

	return session, nil
}

func GetSessionByUserID(userID string) (models.Session, error) {
	session := models.Session{}
	sessionResult := db.GormDB.Where("user_id = ?", userID).First(&session)

	if sessionResult.Error == nil && session.ID != "" {
		return session, errors.New("No session found")
	}

	return session, nil
}

func DeleteSession(session models.Session) error {
	deleteResult := db.GormDB.Delete(&session)

	return deleteResult.Error
}

func CreateSession(userID string) (models.Session, error) {
	session := models.Session{}

	expDate := time.Now().UTC().Add(time.Hour * 24)
	token, err := helpers.CreateToken(userID, expDate)

	if err != nil {
		log.Printf("Error creating token: %v", err)
		return session, errors.New("Internal Server Error")
	}

	session = models.Session{
		ID:        helpers.GetUUID(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expDate,
	}

	createSessionResult := db.GormDB.Create(&session)
	if createSessionResult.Error != nil {
		log.Printf("Error creating session: %v", createSessionResult.Error)
		return session, errors.New("Internal Server Error")
	}

	return session, nil
}

func GetOrCreateSession(userID string) (models.Session, error) {
	session, err := GetSessionByUserID(userID)
	isSessionExpired := isSessionExpired(session)

	if err == nil || isSessionExpired {
		if isSessionExpired {
			err := DeleteSession(session)

			if err != nil {
				log.Printf("Error deleting expired session: %v", err)
				return session, errors.New("Internal Server Error")
			}
		}

		return CreateSession(userID)
	}

	return session, nil
}
