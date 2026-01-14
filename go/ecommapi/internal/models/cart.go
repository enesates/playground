package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID          string  `gorm:"type:varchar(255);primaryKey"`
	UserID      string  `gorm:"type:varchar(255);not null;uniqueindex"`
	TotalAmount float64 `gorm:"type:decimal(10,2);not null;column:total_amount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CartItems []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
}
