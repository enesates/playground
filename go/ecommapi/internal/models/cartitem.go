package models

import (
	"time"
)

type CartItem struct {
	CartID    string `gorm:"type:uuid;primaryKey"`
	ProductID string `gorm:"type:uuid;primaryKey"`
	Quantity  int    `gorm:"type:integer;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (CartItem) TableName() string {
	return "cart_items"
}
