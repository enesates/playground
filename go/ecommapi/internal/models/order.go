package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              string  `gorm:"type:varchar(255);primaryKey"`
	UserID          string  `gorm:"type:varchar(255);not null;index"`
	Status          string  `gorm:"type:varchar(50);not null"`
	TotalAmount     float64 `gorm:"type:decimal(10,2);not null;column:total_amount"`
	ShippingStreet  string  `gorm:"type:varchar(255);not null;column:shipping_street"`
	ShippingCity    string  `gorm:"type:varchar(100);not null;column:shipping_city"`
	ShippingZip     string  `gorm:"type:varchar(20);not null;column:shipping_zip"`
	ShippingCountry string  `gorm:"type:varchar(100);not null;column:shipping_country"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	User       User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}
