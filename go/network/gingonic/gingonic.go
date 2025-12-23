package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Msg string `json:"msg"`
}

func main() {
	router := gin.Default()

	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": time.Now(),
		})
	})

	router.POST("/echo", func(c *gin.Context) {
		var ur UserRequest

		if err := c.ShouldBindJSON(&ur); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(200, ur)
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
