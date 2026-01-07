package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Profile struct {
	gorm.Model
	ID     int
	UserID int
	Bio    string
	Avatar string
	User   User
}
