package db

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
}

type Session struct {
	ID        string    `gorm:"type:varchar(255);primaryKey"`
	UserID    string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null;column:expires_at"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Product struct {
	ID          string  `gorm:"type:varchar(255);primaryKey"`
	Name        string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	CategoryID  string  `gorm:"type:varchar(100);not null;column:category_id"`
	IsActive    bool    `gorm:"type:boolean;not null;default:true;column:is_active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Stock struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	ProductID string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Quantity  int    `gorm:"type:integer;not null"`
	Reserved  int    `gorm:"type:integer;not null"`
	Location  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

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

type CartItem struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	CartID    string `gorm:"type:varchar(255);not null;index;uniqueIndex:idx_cart_product"`
	ProductID string `gorm:"type:varchar(255);not null;uniqueIndex:idx_cart_product"`
	Quantity  int    `gorm:"type:integer;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

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

type OrderItem struct {
	ID        string  `gorm:"type:varchar(255);primaryKey"`
	OrderID   string  `gorm:"type:varchar(255);not null;index"`
	ProductID string  `gorm:"type:varchar(255);not null;index"`
	Quantity  int     `gorm:"type:integer;not null"`
	UnitPrice float64 `gorm:"type:decimal(10,2);not null;column:unit_price"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
}

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
