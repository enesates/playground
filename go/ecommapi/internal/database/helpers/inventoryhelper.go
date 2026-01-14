package db

import (
	db "ecommapi/internal/database"
	"ecommapi/internal/dtos"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"

	"gorm.io/gorm"
)

func GetInventory(pid string) (*models.Stock, error) {
	stock := models.Stock{}
	stockResult := db.GormDB.Where("product_id = ?", pid).First(&stock)

	if stockResult.Error != nil || stockResult.RowsAffected == 0 {
		return nil, stockResult.Error
	}

	return &stock, nil
}

func CreateInventory(pid string, stockDTO dtos.StockDTO) (*models.Stock, error) {
	stock := models.Stock{
		ID:        helpers.GetUUID(),
		ProductID: pid,
		Quantity:  stockDTO.IncerementBy,
		Reserved:  0,
		Location:  "Berlin",
	}

	if err := db.GormDB.Create(&stock).Error; err != nil {
		return nil, err
	}

	return &stock, nil
}

func UpdateInventory(pid string, stockDTO dtos.StockDTO) (*models.Stock, error) {
	tx := db.GormDB.Begin()

	if err := tx.
		Model(&models.Stock{}).
		Where("product_id = ?", pid).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", stockDTO.IncerementBy)).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	stock := models.Stock{}

	if err := tx.
		Where("product_id = ?", pid).
		First(&stock).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &stock, nil
}
