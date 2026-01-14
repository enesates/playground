package db

import (
	"time"

	db "ecommapi/internal/database"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func GetSessionByToken(token string) (*models.Session, error) {
	session := models.Session{
		Token: token,
	}

	if err := db.GormDB.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func GetSessionByUserID(userID string) (*models.Session, error) {
	session := models.Session{}

	if err := db.GormDB.Where("user_id = ?", userID).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func DeleteSession(session models.Session) error {
	deleteResult := db.GormDB.Delete(&session)

	return deleteResult.Error
}

func CreateSession(userID string) (*models.Session, error) {
	expDate := time.Now().UTC().Add(time.Hour * 24)
	token, err := helpers.CreateToken(userID, expDate)

	if err != nil {
		return nil, err
	}

	session := models.Session{
		ID:        helpers.GetUUID(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expDate,
	}

	if err := db.GormDB.Create(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func GetOrCreateSession(userID string) (*models.Session, error) {
	session, _ := GetSessionByUserID(userID)

	if session == nil {
		return CreateSession(userID)
	}

	isSessionValid := isSessionValid(*session)
	isSessionExpired := isSessionExpired(*session)

	if !isSessionValid || isSessionExpired {
		if err := DeleteSession(*session); err != nil {
			return nil, err
		}

		return CreateSession(userID)
	}

	return session, nil
}
