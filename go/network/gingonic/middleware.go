package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserRequest2 struct {
	Msg string `json:"msg"`
}

func checkHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-App-Version") != "1.0" {
			c.String(http.StatusInternalServerError, "INTERNAL SERVER ERROR")
			c.Abort()
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(checkHeaderMiddleware())

	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": time.Now(),
		})
	})

	router.POST("/echo", func(c *gin.Context) {
		var ur UserRequest2

		if err := c.ShouldBindJSON(&ur); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(200, ur)
	})

	router.Run()
}
