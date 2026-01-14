package handlers

import (
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/dtos"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Create a User
// @Description Creating a new Customer
// @Tags user
// @Accept json
// @Produce json
// @Param data body dtos.UserRegisterDTO true "New User"
// @Success 200 {object} map[string]any "Details of the User"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var userDTO dtos.UserRegisterDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(err))
		return
	}

	if dbHelper.CheckUserExists(userDTO.Username, userDTO.Email) {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(errors.New("User already exists")))
		return
	}

	user, err := dbHelper.CreateUser(userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

// Login godoc
// @Summary Login
// @Description Login as Customer or Admin
// @Tags user
// @Accept json
// @Produce json
// @Param data body dtos.UserLoginDTO true "Login credentials"
// @Success 200 {object} map[string]any "Details of the User"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var userDTO dtos.UserLoginDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, session, err := dbHelper.GetUserAndSession(userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sesion_token": session.Token,
		"expires_at":   session.ExpiresAt,
		"user": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

// Logout godoc
// @Summary Logout
// @Description Logging out as a User
// @Tags user
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Success 200 {object} map[string]any "Details of the User"
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	token := c.GetHeader("X-Session-Token")
	session, err := dbHelper.GetSessionByToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	user, err := dbHelper.GetUserByID(session.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	err = dbHelper.DeleteSession(*session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"sesion_token": session.Token,
		"expires_at":   session.ExpiresAt,
		"user": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
