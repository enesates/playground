package main

import (
	"fmt"
	"net/http"
	"strings"
	"webserver/internal/http/handlers"

	"github.com/stretchr/testify/assert"

	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	router := handlers.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	messageJson := `{"message":"OK"}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, messageJson, w.Body.String())
}

func TestEchoHandler(t *testing.T) {
	t.Run("Proper message", func(t *testing.T) {
		router := handlers.SetupRouter()
		w := httptest.NewRecorder()
		messageText := "Hello Gin!"

		path := strings.Join([]string{"/echo?message=", messageText}, "")
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		messageJson := fmt.Sprintf(`{"message":"%s"}`, messageText)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, messageJson, w.Body.String())
	})

	t.Run("Empty message", func(t *testing.T) {
		router := handlers.SetupRouter()
		w := httptest.NewRecorder()

		path := "/echo?message="
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		messageJson := `{"message":"Missing Message"}`

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, messageJson, w.Body.String())
	})

}
