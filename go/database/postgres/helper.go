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

        userId, err := AddUser(db, username, email)
        if err != nil || userId == -1 {
            log.Printf("Error while creating user %d\n", userId)
        }
        log.Printf("User %d created\n", userId)
    }
}
