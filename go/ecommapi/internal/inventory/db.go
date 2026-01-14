package inventory

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"

	"gorm.io/gorm"
)

func FetchInventory(pid string) (*db.Stock, error) {
	stock := db.Stock{}
	stockResult := db.GormDB.Where("product_id = ?", pid).First(&stock)

	if stockResult.Error != nil || stockResult.RowsAffected == 0 {
		return nil, stockResult.Error
	}

	return &stock, nil
}

func CreateInventory(pid string, stockDTO StockDTO) (*db.Stock, error) {
	stock := db.Stock{
		ID:        utils.GetUUID(),
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

func UpdateInventory(pid string, stockDTO StockDTO) (*db.Stock, error) {
	tx := db.GormDB.Begin()

	if err := tx.
		Model(&db.Stock{}).
		Where("product_id = ?", pid).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", stockDTO.IncerementBy)).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	stock := db.Stock{}

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
