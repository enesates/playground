package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func init() {
	var url = "postgres://postgresuser:postgrespass@localhost:5432/mydb?sslmode=disable"
	PostgresDB, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err = PostgresDB.Ping(); err != nil {
		log.Fatalf("Error on DB connection: %v", err)
	}

	fmt.Println("DB connected!")
}

func main() {
	defer closeConnection(PostgresDB)
	CreateUsersTable()
}
