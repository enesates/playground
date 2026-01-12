package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, email string, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("User couldn't created. Error:", err)
		return
	}
	user := User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	AddUser(user)
}

// if static salt necessary
// salt, _ := bcrypt.Salt()
// hash, _ = bcrypt.Hash(password, salt)
// if bcrypt.Match(password, hash) {
// 	fmt.Println("They match")
// }
