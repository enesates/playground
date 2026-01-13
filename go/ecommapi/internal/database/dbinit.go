package db

import (
	"ecommapi/internal/models"
	"log"

	"github.com/gofor-little/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func CreateTables() {
	err := GormDB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Order{},
		&models.Product{},
		&models.Stock{},
		&models.OrderItem{},
		&models.CartItem{},
		&models.Notification{},
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
