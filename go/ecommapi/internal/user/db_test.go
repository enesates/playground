package user

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUserExists(t *testing.T) {
	db.SetupTestDB()

	u := db.User{
		ID:           "test-user-id",
		Username:     "testuser",
		Email:        "testuser@ecommapi.com",
		PasswordHash: "password",
	}

	db.GormDB.Create(&u)

	exists := CheckUserExists("testuser", "testuser@ecommapi.com")
	assert.True(t, exists)

	notExists := CheckUserExists("x", "y@test.com")
	assert.False(t, notExists)
}

func TestCreateUser(t *testing.T) {
	db.SetupTestDB()

	userDTO := UserRegisterDTO{
		Username: "testuser",
		Email:    "testuser@ecommapi.com",
		Password: "password",
	}

	user, err := CreateUser(userDTO)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, userDTO.Username, user.Username)

	user, err = CreateUser(userDTO)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestGetUserAndSession_Success(t *testing.T) {
	db.SetupTestDB()

	password := "password"
	hash, _ := utils.HashPassword(password)

	user := db.User{
		ID:           "test-user-id",
		Username:     "testuser",
		Email:        "testuser@ecommapi.com",
		PasswordHash: hash,
	}
	db.GormDB.Create(&user)

	dto := UserLoginDTO{
		Email:    "testuser@ecommapi.com",
		Password: "password",
	}
	u, session, err := GetUserAndSession(dto)

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotNil(t, session)
}

func TestGetUserAndSession_InvalidPassword(t *testing.T) {
	db.SetupTestDB()

	user := db.User{
		ID:           "test-user-id",
		Username:     "testuser",
		Email:        "testuser@ecommapi.com",
		PasswordHash: "password",
	}
	db.GormDB.Create(&user)

	dto := UserLoginDTO{
		Email:    "testuser@ecommapi.com",
		Password: "password",
	}
	u, s, err := GetUserAndSession(dto)

	assert.Error(t, err)
	assert.Nil(t, u)
	assert.Nil(t, s)
}

func TestGetUserByID_Success(t *testing.T) {
	db.SetupTestDB()

	user := db.User{
		ID:           "test-user-id",
		Username:     "testuser",
		Email:        "testuser@ecommapi.com",
		PasswordHash: "password",
	}
	db.GormDB.Create(&user)

	u, err := GetUserByID("test-user-id")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", u.Username)
}

func TestGetUserByID_NotFound(t *testing.T) {
	db.SetupTestDB()

	user := db.User{
		ID:           "test-user-id",
		Username:     "testuser",
		Email:        "testuser@ecommapi.com",
		PasswordHash: "password",
	}
	db.GormDB.Create(&user)

	u, err := GetUserByID("wrong-id")

	assert.Error(t, err)
	assert.Nil(t, u)
}
