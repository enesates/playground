package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string, expDate time.Time) (string, error) {
	var secretKey = []byte("secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userID,
			"exp":      expDate,
		})

	return token.SignedString(secretKey)
}
