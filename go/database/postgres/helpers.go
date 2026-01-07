package main

import (
	"fmt"
	"log"
)

func CloseConnection() {
	defer func() {
		err := PostgresDB.Close()
		if err != nil {
			log.Fatalf("Error on closing DB connection: %v", err)
		}
	}()
}

func CreateUsers() {
	for i := 1; i < 10; i++ {
		username := fmt.Sprintf("NewUser%d", i)
		email := fmt.Sprintf("new@user%d.com", i)
		user := User{name: username, email: email}

		userId, err := AddUser(user)
		if err != nil || userId == -1 {
			log.Printf("Error while creating user %d\n", userId)
		}
		log.Printf("User %d created\n", userId)
	}
}

func GetAllUsers() {
	users, err := GetUsers()
	if err != nil {
		log.Fatal("Error getting users: ", err)
	}

	log.Printf("\n\nFound %d users\n----------\n", len(users))
	for _, user := range users {
		log.Printf("User %s\n", user.name)
	}
}
