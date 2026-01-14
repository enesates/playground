package models

import (
	"time"
)

type CartItem struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	CartID    string `gorm:"type:varchar(255);not null;index;uniqueIndex:idx_cart_product"`
	ProductID string `gorm:"type:varchar(255);not null;uniqueIndex:idx_cart_product"`
	Quantity  int    `gorm:"type:integer;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
