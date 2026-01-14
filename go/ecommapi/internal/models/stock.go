package models

import (
	"time"
)

type Stock struct {
	ProductID string `gorm:"type:varchar(255);primaryKey"`
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
