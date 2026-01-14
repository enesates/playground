package models

import (
	"time"
)

type Stock struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	ProductID string `gorm:"type:varchar(255);not null;uniqueIndex"`
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
