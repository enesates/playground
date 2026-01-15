package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.POST("/health", Health)

	return r
}

func TestHealth(t *testing.T) {
	router := setupRouter()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/health", nil)
		// req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
