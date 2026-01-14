package db

import (
	db "ecommapi/internal/database"
	"ecommapi/internal/dtos"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"
)

func GetProductByID(pid string) (*models.Product, error) {
	product := models.Product{}

	if err := db.GormDB.Where("id = ?", pid).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProducts(cid string, page int) ([]models.Product, error) {
	products := []models.Product{}

	if err := db.GormDB.
		Limit(10).Offset((page-1)*10).
		Where("category_id = ?", cid).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(productDTO dtos.ProductDTO) (*models.Product, error) {
	product := models.Product{
		ID:          helpers.GetUUID(),
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		CategoryID:  productDTO.CategoryID,
		IsActive:    true,
	}

	if err := db.GormDB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
