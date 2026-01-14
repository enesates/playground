package api

import (
	"net/http"

	"ecommapi/internal/auth"
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/user"

	"github.com/gin-gonic/gin"
)

func isTokenValid(token string) bool {
	if token == "" {
		return false
	}

	_, err := auth.GetSessionByToken(token)

	return err == nil
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Session-Token")

		if !isTokenValid(token) {
			utils.AbortJSON(c, http.StatusUnauthorized, "Invalid token")
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
		utils.AbortJSON(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	user, err := user.GetUserByToken(token)

	if err != nil || user.Role != role {
		utils.AbortJSON(c, http.StatusForbidden, "Authorization failed")
		return
	}
}
