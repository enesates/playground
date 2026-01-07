package main

import (
	"database/sql"
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
