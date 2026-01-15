package user

import (
	"bytes"
	db "ecommapi/internal/helpers/database"
	notif "ecommapi/internal/notification"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/auth/register", Register)
	return r
}

func performRequest(router *gin.Engine, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestRegister(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name           string
		mockCheckUser  func(username, email string) bool
		mockCreateUser func(dto UserRegisterDTO) (*db.User, error)
		mockNotif      func(uid, title, msg string) error
		requestBody    string
		expectedStatus int
	}{
		{
			name:          "success",
			mockCheckUser: func(_, _ string) bool { return false },
			mockCreateUser: func(dto UserRegisterDTO) (*db.User, error) {
				return &db.User{ID: "test-user-id", Username: "usertest", Email: "usertest@ecommapi.com"}, nil
			},
			mockNotif:      func(_, _, _ string) error { return nil },
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid request",
			mockCheckUser:  nil,
			mockCreateUser: nil,
			mockNotif:      nil,
			requestBody:    `{ invalid request }`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "user exists",
			mockCheckUser:  func(_, _ string) bool { return true },
			mockCreateUser: nil,
			mockNotif:      nil,
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "user creation fails",
			mockCheckUser:  func(_, _ string) bool { return false },
			mockCreateUser: func(dto UserRegisterDTO) (*db.User, error) { return nil, errors.New("error") },
			mockNotif:      nil,
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:          "notification fails",
			mockCheckUser: func(_, _ string) bool { return false },
			mockCreateUser: func(dto UserRegisterDTO) (*db.User, error) {
				return &db.User{ID: "test-user-id", Username: "usertest", Email: "usertest@ecommapi.com"}, nil
			},
			mockNotif:      func(uid, title, msg string) error { return errors.New("error") },
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockCheckUser != nil {
				checkUserExists = tt.mockCheckUser
				defer func() { checkUserExists = CheckUserExists }()
			}
			if tt.mockCreateUser != nil {
				createUser = tt.mockCreateUser
				defer func() { createUser = CreateUser }()
			}
			if tt.mockNotif != nil {
				createEventNotif = tt.mockNotif
				defer func() { createEventNotif = notif.CreateEventNotif }()
			}

			w := performRequest(router, tt.requestBody)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
