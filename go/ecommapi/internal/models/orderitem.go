package models

import (
	"time"
)

type OrderItem struct {
	OrderID   string  `gorm:"type:varchar(255);primaryKey"`
	ProductID string  `gorm:"type:varchar(255);primaryKey"`
	Quantity  int     `gorm:"type:integer;not null"`
	UnitPrice float64 `gorm:"type:decimal(10,2);not null;column:unit_price"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
