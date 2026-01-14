package helpers

import (
	"github.com/gin-gonic/gin"
)

func AbortJSON(c *gin.Context, status int, msg string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": msg,
	})
}
