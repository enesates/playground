package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Health Check
// @Description Check if service runs properly
// @Tags health
// @Produce json
// @Success 200 {object} map[string]any "Health message"
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
