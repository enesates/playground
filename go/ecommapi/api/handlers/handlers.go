package handlers

import (
	db "ecommapi/internal/database"
	"ecommapi/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func RegisterUser(c *gin.Context) {
	var userDTO models.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if db.CheckUserExists(userDTO) {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(errors.New("User already exists")))
		return
	}

	err := db.CreateUser(userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	// c.JSON(200, user)
}
