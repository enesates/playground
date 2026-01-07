package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func init() {
	var err error

	dsn := "host=localhost user=postgresuser password=postgrespass dbname=mydb port=5432 sslmode=disable"
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	CreateTables()

	CreateUsers()

	GetAllUsers()
}
