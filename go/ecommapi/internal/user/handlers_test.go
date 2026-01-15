package user

import (
	"bytes"
	"ecommapi/internal/auth"
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
	r.POST("/auth/login", Login)
	r.POST("/auth/logout", Logout)

	return r
}

func performRequest(router *gin.Engine, method string, path string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
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
			mockCreateUser: func(_ UserRegisterDTO) (*db.User, error) {
				return &db.User{}, nil
			},
			mockNotif:      func(_, _, _ string) error { return nil },
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid request",
			requestBody:    `{ invalid request }`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "user exists",
			mockCheckUser:  func(_, _ string) bool { return true },
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "user creation fails",
			mockCheckUser:  func(_, _ string) bool { return false },
			mockCreateUser: func(_ UserRegisterDTO) (*db.User, error) { return nil, errors.New("error") },
			requestBody:    `{"username":"usertest","email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "notification fails",
			mockCheckUser:  func(_, _ string) bool { return false },
			mockCreateUser: func(_ UserRegisterDTO) (*db.User, error) { return &db.User{}, nil },
			mockNotif:      func(_, _, _ string) error { return errors.New("error") },
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

			w := performRequest(router, http.MethodPost, "/auth/register", tt.requestBody)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestLogin(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name               string
		mockGetUserSession func(_ UserLoginDTO) (*db.User, *db.Session, error)
		mockNotif          func(_, _, _ string) error
		requestBody        string
		expectedStatus     int
	}{
		{
			name:               "success",
			mockGetUserSession: func(_ UserLoginDTO) (*db.User, *db.Session, error) { return &db.User{}, &db.Session{}, nil },
			mockNotif:          func(_, _, _ string) error { return nil },
			requestBody:        `{"email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus:     http.StatusOK,
		},
		{
			name:           "invalid request",
			requestBody:    `{ invalid request }`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:               "get user and session fails",
			mockGetUserSession: func(_ UserLoginDTO) (*db.User, *db.Session, error) { return nil, nil, errors.New("error") },
			requestBody:        `{"email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus:     http.StatusInternalServerError,
		},
		{
			name:               "notification fails",
			mockGetUserSession: func(_ UserLoginDTO) (*db.User, *db.Session, error) { return &db.User{}, &db.Session{}, nil },
			mockNotif:          func(_, _, _ string) error { return errors.New("error") },
			requestBody:        `{"email":"usertest@ecommapi.com","password":"usertest"}`,
			expectedStatus:     http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockGetUserSession != nil {
				getUserAndSession = tt.mockGetUserSession
				defer func() { getUserAndSession = GetUserAndSession }()
			}
			if tt.mockNotif != nil {
				createEventNotif = tt.mockNotif
				defer func() { createEventNotif = notif.CreateEventNotif }()
			}

			w := performRequest(router, http.MethodPost, "/auth/login", tt.requestBody)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestLogout(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name                  string
		mockGetSessionByToken func(_ string) (*db.Session, error)
		mockGetUserByID       func(_ string) (*db.User, error)
		mockDeleteSession     func(_ db.Session) error
		mockNotif             func(_, _, _ string) error
		expectedStatus        int
	}{
		{
			name:                  "success",
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockGetUserByID:       func(_ string) (*db.User, error) { return &db.User{}, nil },
			mockDeleteSession:     func(_ db.Session) error { return nil },
			mockNotif:             func(_, _, _ string) error { return nil },
			expectedStatus:        http.StatusNoContent,
		},
		{
			name:                  "get session fails",
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return nil, errors.New("error") },
			expectedStatus:        http.StatusInternalServerError,
		},
		{
			name:                  "get user fails",
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockGetUserByID:       func(_ string) (*db.User, error) { return nil, errors.New("error") },
			expectedStatus:        http.StatusInternalServerError,
		},
		{
			name:                  "delete session fails",
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockGetUserByID:       func(_ string) (*db.User, error) { return &db.User{}, nil },
			mockDeleteSession:     func(_ db.Session) error { return errors.New("error") },
			expectedStatus:        http.StatusInternalServerError,
		},
		{
			name:                  "notification fails",
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockGetUserByID:       func(_ string) (*db.User, error) { return &db.User{}, nil },
			mockDeleteSession:     func(_ db.Session) error { return nil },
			mockNotif:             func(_, _, _ string) error { return errors.New("error") },
			expectedStatus:        http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockGetSessionByToken != nil {
				getSessionByToken = tt.mockGetSessionByToken
				defer func() { getSessionByToken = auth.GetSessionByToken }()
			}
			if tt.mockGetUserByID != nil {
				getUserByID = tt.mockGetUserByID
				defer func() { getUserByID = GetUserByID }()
			}
			if tt.mockDeleteSession != nil {
				deleteSession = tt.mockDeleteSession
				defer func() { deleteSession = auth.DeleteSession }()
			}
			if tt.mockNotif != nil {
				createEventNotif = tt.mockNotif
				defer func() { createEventNotif = notif.CreateEventNotif }()
			}

			w := performRequest(router, http.MethodPost, "/auth/logout", "")
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
