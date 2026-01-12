package database

import (
	"time"

	"gorm.io/gorm"
)

// User represents the users table
type User struct {
	ID           string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username     string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Email        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	PasswordHash string `gorm:"type:text;not null;column:password_hash"`
	Role         string `gorm:"type:varchar(50)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// Relationships
	Sessions      []Session      `gorm:"foreignKey:UserID"`
	Orders        []Order        `gorm:"foreignKey:UserID"`
	Notifications []Notification `gorm:"foreignKey:UserID"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

// Session represents the sessions table
type Session struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    string    `gorm:"type:uuid;not null;index"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null;column:expires_at"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for Session model
func (Session) TableName() string {
	return "sessions"
}

// Order represents the orders table
type Order struct {
	ID              string  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID          string  `gorm:"type:uuid;not null;index"`
	Status          string  `gorm:"type:varchar(50);not null"`
	TotalAmount     float64 `gorm:"type:decimal(10,2);not null;column:total_amount"`
	ShippingStreet  string  `gorm:"type:varchar(255);not null;column:shipping_street"`
	ShippingCity    string  `gorm:"type:varchar(100);not null;column:shipping_city"`
	ShippingZip     string  `gorm:"type:varchar(20);not null;column:shipping_zip"`
	ShippingCountry string  `gorm:"type:varchar(100);not null;column:shipping_country"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	// Relationships
	User       User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for Order model
func (Order) TableName() string {
	return "orders"
}

// Product represents the products table
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

	// Relationships
	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID"`
}

// TableName specifies the table name for Product model
func (Product) TableName() string {
	return "products"
}

// Stock represents the stocks table
type Stock struct {
	ProductID string `gorm:"type:uuid;primaryKey"`
	Quantity  int    `gorm:"type:integer;not null"`
	Reserved  bool   `gorm:"type:boolean;not null;default:false"`
	Location  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for Stock model
func (Stock) TableName() string {
	return "stocks"
}

// OrderItem represents the order_items table
type OrderItem struct {
	OrderID   string  `gorm:"type:uuid;primaryKey"`
	ProductID string  `gorm:"type:uuid;primaryKey"`
	Quantity  int     `gorm:"type:integer;not null"`
	UnitPrice float64 `gorm:"type:decimal(10,2);not null;column:unit_price"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
}

// TableName specifies the table name for OrderItem model
func (OrderItem) TableName() string {
	return "order_items"
}

// CartItem represents the cart_items table
type CartItem struct {
	CartID    string `gorm:"type:uuid;primaryKey"`
	ProductID string `gorm:"type:uuid;primaryKey"`
	Quantity  int    `gorm:"type:integer;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relationships
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for CartItem model
func (CartItem) TableName() string {
	return "cart_items"
}

// Notification represents the notifications table
type Notification struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    string `gorm:"type:uuid;not null;index"`
	Title     string `gorm:"type:varchar(255);not null"`
	Message   string `gorm:"type:text;not null"`
	IsRead    bool   `gorm:"type:boolean;not null;default:false;column:is_read"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for Notification model
func (Notification) TableName() string {
	return "notifications"
}
