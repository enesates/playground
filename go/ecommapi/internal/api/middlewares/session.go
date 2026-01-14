package middlewares

import (
	"net/http"

	h "ecommapi/internal/api/helpers"
	dbHelper "ecommapi/internal/database/helpers"

	"github.com/gin-gonic/gin"
)

func isTokenValid(token string) bool {
	if token == "" {
		return false
	}

	_, err := dbHelper.GetSessionByToken(token)

	return err == nil
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Session-Token")

		if !isTokenValid(token) {
			h.AbortJSON(c, http.StatusUnauthorized, "Invalid token")
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
		h.AbortJSON(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	user, err := dbHelper.GetUserByToken(token)

	if err != nil || user.Role != role {
		h.AbortJSON(c, http.StatusForbidden, "Authorization failed")
		c.Abort()
		return
	}
}
