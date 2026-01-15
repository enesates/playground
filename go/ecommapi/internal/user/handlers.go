package user

import (
	"ecommapi/internal/auth"
	"ecommapi/internal/helpers/utils"
	notif "ecommapi/internal/notification"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Create a User
// @Description Creating a new Customer
// @Tags user
// @Accept json
// @Produce json
// @Param data body UserRegisterDTO true "New User"
// @Success 200 {object} map[string]any "Details of the User"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var userDTO UserRegisterDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if CheckUserExists(userDTO.Username, userDTO.Email) {
		utils.AbortJSON(c, http.StatusBadRequest, "User already exists")
		return
	}

	user, err := CreateUser(userDTO)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notif.CreateEventNotif(user.ID, "Register", "User registered"); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
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
// @Param data body UserLoginDTO true "Login credentials"
// @Success 200 {object} map[string]any "Details of the User"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var userDTO UserLoginDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, session, err := GetUserAndSession(userDTO)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notif.CreateEventNotif(user.ID, "Login", "User logged in"); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
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
	session, err := auth.GetSessionByToken(token)

	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := GetUserByID(session.UserID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = auth.DeleteSession(*session)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notif.CreateEventNotif(user.ID, "Logout", "User logged out"); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
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
