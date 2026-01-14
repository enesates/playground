package db

import (
	"time"

	"ecommapi/internal/models"
)

func isSessionExpired(session models.Session) bool {
	if session.ExpiresAt.IsZero() {
		return false
	}

	now := time.Now().UTC()
	return session.ExpiresAt.UTC().Before(now)
}

func isSessionValid(session models.Session) bool {
	return session.Token != "" && session.UserID != "" && !session.ExpiresAt.IsZero()
}
