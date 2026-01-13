package middlewares

import (
	"errors"
	"net/http"

	dbHelper "ecommapi/internal/database/helpers"

	"github.com/gin-gonic/gin"
)

func CheckSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken := c.GetHeader("X-Session-Token")

		if sessionToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, c.Error(errors.New("Invalid Token")))
			c.Abort()
			return
		}

		_, err := dbHelper.GetSessionByToken(sessionToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, c.Error(errors.New("Invalid Token")))
			c.Abort()
			return
		}

		c.Next()
	}
}
