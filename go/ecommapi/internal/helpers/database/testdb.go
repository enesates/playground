package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB() {
	var err error
	GormDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

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
