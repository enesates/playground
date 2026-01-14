package db

import (
	"log"

	"github.com/gofor-little/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func DBInit() {
	var err error

	if err = env.Load(".env"); err != nil {
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

	err = GormDB.AutoMigrate(
		&User{},
		&Session{},
		&Product{},
		&Stock{},
		&Cart{},
		&CartItem{},
		&Order{},
		&OrderItem{},
		&Notification{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
