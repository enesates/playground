package models

import (
	"time"
)

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
