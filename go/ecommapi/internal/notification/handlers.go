package notif

import (
	"ecommapi/internal/auth"
	"ecommapi/internal/helpers/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotifications godoc
// @Summary Get notifications
// @Description Get notifications of a user
// @Tags notification
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Success 200 {object} map[string]any "Notifications"
// @Router /notifications [get]
func GetNotifications(c *gin.Context) {
	token := c.GetHeader("X-Session-Token")
	session, err := auth.GetSessionByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	notifications, err := FetchNotifications(session.User.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	nItems := []gin.H{}

	for _, n := range notifications {
		nItems = append(nItems, gin.H{
			"username": n.User.Username,
			"title":    n.Title,
			"message":  n.Message,
			"is_read":  n.IsRead,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"notifications": nItems,
	})
}

// MakeNotificationRead godoc
// @Summary Update notification status
// @Description Update notification read status true
// @Tags notification
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param id path string true "Notification ID"
// @Success 200 {object} map[string]any "Notification data"
// @Router /notifications/:id/read [patch]
func MakeNotificationRead(c *gin.Context) {
	nid := c.Param("id")

	if err := UpdateNotificationStatus(nid); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	notification, err := FetchNotification(nid)

	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notification": gin.H{
			"username": notification.User.Username,
			"title":    notification.Title,
			"message":  notification.Message,
			"is_read":  notification.IsRead,
		},
	})
}
