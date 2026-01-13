package middlewares

import (
	"errors"
	"net/http"

	dbHelper "ecommapi/internal/database/helpers"

	"github.com/gin-gonic/gin"
)

func isTokenValid(token string) bool {
	if token == "" {
		return false
	}

	_, err := dbHelper.GetSessionByToken(token)

	if err != nil {
		return false
	}

	return true
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Session-Token")

		if !isTokenValid(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, c.Error(errors.New("Invalid Token")))
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckIsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		CheckAuthorizationForRole(c, "admin")
		c.Next()
	}
}

func CheckIsCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		CheckAuthorizationForRole(c, "customer")
		c.Next()
	}
}

func CheckAuthorizationForRole(c *gin.Context, role string) {
	token := c.GetHeader("X-Session-Token")

	if !isTokenValid(token) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, c.Error(errors.New("Invalid Token")))
		c.Abort()
		return
	}

	user, err := dbHelper.GetUserByToken(token)

	if err != nil || user.Role != role {
		c.AbortWithStatusJSON(http.StatusForbidden, c.Error(errors.New("Authorization failed")))
		c.Abort()
		return
	}
}
