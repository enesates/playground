package main

import (
	"database/sql"
	"fmt"
	"log"
)

func CloseConnection(db *sql.DB) {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error on closing DB connection: %v", err)
		}
	}()
}

func CreateUsers(db *sql.DB) {
	for i := 1; i < 10; i++ {
		username := fmt.Sprintf("NewUser%d", i)
		email := fmt.Sprintf("new@user%d.com", i)
		user := User{name: username, email: email}

		userId, err := AddUser(db, user)
		if err != nil || userId == -1 {
			log.Printf("Error while creating user %d\n", userId)
		}
		log.Printf("User %d created\n", userId)
	}
}

func GetAllUsers(db *sql.DB) {
	users, err := GetUsers(db)
	if err != nil {
		log.Fatal("Error getting users: ", err)
	}

	log.Printf("\n\nFound %d users\n----------\n", len(users))
	for _, user := range users {
		log.Printf("User %s\n", user.name)
	}
}
