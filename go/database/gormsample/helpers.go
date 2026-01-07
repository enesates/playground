package main

import (
	"fmt"
	"log"
)

func CreateUsers() {
	for i := 1; i < 10; i++ {
		username := fmt.Sprintf("NewUser%d", i)
		email := fmt.Sprintf("new@user%d.com", i)
		user := User{Name: username, Email: email, Age: i * 10}

		AddUser(user)
	}
}

func GetAllUsers() {
	users := GetUsers()

	log.Printf("\n\nFound %d users\n----------\n", len(users))

	for _, user := range users {
		log.Printf("User %v\n", user)
	}
}
