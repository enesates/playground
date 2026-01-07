package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func CreateUsersTable(db *sql.DB) {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	fmt.Println("Created users table")
}

func AddUser(db *sql.DB, username string, email string) (int, error) {
	var id int

	err := db.QueryRow(`
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, username, email, time.Now()).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
