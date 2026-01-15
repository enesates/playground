package auth

import (
	"time"

	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
)

func GetSessionByToken(token string) (*db.Session, error) {
	session := db.Session{
		Token: token,
	}

	if err := db.GormDB.Preload("User").Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func GetSessionByUserID(userID string) (*db.Session, error) {
	session := db.Session{}

	if err := db.GormDB.Where("user_id = ?", userID).First(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func DeleteSession(session db.Session) error {
	deleteResult := db.GormDB.Delete(&session)

	return deleteResult.Error
}

func CreateSession(userID string) (*db.Session, error) {
	expDate := time.Now().UTC().Add(time.Hour * 24)
	token, err := CreateToken(userID, expDate)

	if err != nil {
		return nil, err
	}

	session := db.Session{
		ID:        utils.GetUUID(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expDate,
	}

	if err := db.GormDB.Create(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}
