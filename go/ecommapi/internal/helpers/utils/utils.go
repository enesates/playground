package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUUID() string {
	return shortuuid.New()
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func AbortJSON(c *gin.Context, status int, msg string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": msg,
	})
}
