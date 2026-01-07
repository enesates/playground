package main

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateUsersTable(db *sql.DB) {
	if _, err := db.Exec(`
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	fmt.Println("Created users table")
}
