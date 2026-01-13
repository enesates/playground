package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"type:varchar(255);primaryKey"`
	Username     string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Email        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	PasswordHash string `gorm:"type:text;not null;column:password_hash"`
	Role         string `gorm:"type:varchar(50)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Sessions      []Session      `gorm:"foreignKey:UserID"`
	Orders        []Order        `gorm:"foreignKey:UserID"`
	Notifications []Notification `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
