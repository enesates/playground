package main

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Profile struct {
	ID     int
	UserID int
	Bio    string
	Avatar string
	User   User
}
