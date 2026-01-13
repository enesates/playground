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
	var userDTO models.UserRegisterDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(err))
		return
	}

	if db.CheckUserExists(userDTO.Username, userDTO.Email) {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(errors.New("User already exists")))
		return
	}

	user, err := db.CreateUser(userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	// Return user data (excluding sensitive fields like PasswordHash)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func LoginUser(c *gin.Context) {
	var userDTO models.UserLoginDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(200, gin.H{
		"sesion_token": "string",
		"expires_at":   "timestamp",
		"user": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})

	// session, err := db.CreateSession(user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
	// 	return
	// }

}
