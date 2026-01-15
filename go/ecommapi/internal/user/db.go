package user

import (
	"errors"

	"ecommapi/internal/auth"
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
)

func CheckUserExists(username string, email string) bool {
	user := db.User{
		Username: username,
		Email:    email,
	}

	result := db.GormDB.Where("username = ? OR email = ?", user.Username, user.Email).First(&user)
	return result.RowsAffected > 0
}

func CreateUser(userDTO UserRegisterDTO) (*db.User, error) {
	hashedPassword, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	user := db.User{
		ID:           utils.GetUUID(),
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: string(hashedPassword),
		Role:         "customer",
	}

	if err := db.GormDB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserAndSession(userDTO UserLoginDTO) (*db.User, *db.Session, error) {
	user := db.User{
		Email: userDTO.Email,
	}

	if err := db.GormDB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return nil, nil, err
	} else if !utils.ComparePassword(user.PasswordHash, userDTO.Password) {
		return nil, nil, errors.New("Invalid credentials")
	}

	session, err := auth.GetOrCreateSession(user.ID)
	return &user, session, err
}

func GetUserByID(userID string) (*db.User, error) {
	user := db.User{}

	if err := db.GormDB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
