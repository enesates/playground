package database

import (
	"log"

	"github.com/gofor-little/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func CreateTables() {
	err := GormDB.AutoMigrate(
		&User{},
		&Session{},
		&Order{},
		&Product{},
		&Stock{},
		&OrderItem{},
		&CartItem{},
		&Notification{},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func DBInit() {
	var err error

	if err = env.Load("../../.env"); err != nil {
		panic(err)
	}

	dsn, err := env.MustGet("dsn")
	if err != nil {
		panic(err)
	}

	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	CreateTables()
}
