package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	UserID    string `gorm:"type:varchar(255);not null;index"`
	Title     string `gorm:"type:varchar(255);not null"`
	Message   string `gorm:"type:text;not null"`
	IsRead    bool   `gorm:"type:boolean;not null;default:false;column:is_read"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Notification) TableName() string {
	return "notifications"
}
