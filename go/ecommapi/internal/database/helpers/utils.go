package db

import (
	"time"

	"ecommapi/internal/models"
)

func isSessionExpired(session models.Session) bool {
	now := time.Now().UTC()

	return session.ID == "" || session.Token == "" || session.ExpiresAt.UTC().Before(now)
}
