package database

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

type Session struct {
	ID        string    `gorm:"type:varchar(255);primaryKey"`
	UserID    string    `gorm:"type:varchar(255);not null;index"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null;column:expires_at"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Session) TableName() string {
	return "sessions"
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

func (Order) TableName() string {
	return "orders"
}

type Product struct {
	ID          string  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
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

type Stock struct {
	ProductID string `gorm:"type:uuid;primaryKey"`
	Quantity  int    `gorm:"type:integer;not null"`
	Reserved  bool   `gorm:"type:boolean;not null;default:false"`
	Location  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (Stock) TableName() string {
	return "stocks"
}

type OrderItem struct {
	OrderID   string  `gorm:"type:uuid;primaryKey"`
	ProductID string  `gorm:"type:uuid;primaryKey"`
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
