package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string  `gorm:"type:varchar(255);primaryKey;default:gen_random_uuid()"`
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	CategoryID  string  `gorm:"type:varchar(100);not null;column:category_id"`
	IsActive    bool    `gorm:"type:boolean;not null;default:true;column:is_active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "products"
}
