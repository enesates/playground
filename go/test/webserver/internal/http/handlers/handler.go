package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", Health)
	router.GET("/echo", Echo)

	return router
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Echo(c *gin.Context) {
	message := c.Query("message")

	if len(message) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing Message",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
