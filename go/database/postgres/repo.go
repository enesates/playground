package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func CreateUsersTable() {
	if _, err := PostgresDB.Exec(`
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

func AddUser(user User) (int, error) {
	var id int

	err := PostgresDB.QueryRow(`
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.name, user.email, time.Now()).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetUsers() ([]User, error) {
	var users []User

	rows, err := PostgresDB.Query(`
		SELECT id, name, email, created_at
		FROM users
	`)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	if err != nil {
		log.Println("Failed to get users table")
		return nil, err
	}

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.id, &user.name, &user.email, &user.createAt); err != nil {
			log.Println("Failed to get users")
		}

		users = append(users, user)
	}

	return users, nil
}
