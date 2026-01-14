package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	db "ecommapi/internal/helpers/database"
)

func isSessionExpired(session db.Session) bool {
	if session.ExpiresAt.IsZero() {
		return false
	}

	now := time.Now().UTC()
	return session.ExpiresAt.UTC().Before(now)
}

func isSessionValid(session db.Session) bool {
	return session.Token != "" && session.UserID != "" && !session.ExpiresAt.IsZero()
}

func CreateToken(userID string, expDate time.Time) (string, error) {
	var secretKey = []byte("secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userID,
			"exp":      expDate,
		})

	return token.SignedString(secretKey)
}

func GetOrCreateSession(userID string) (*db.Session, error) {
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
