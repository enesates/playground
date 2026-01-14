package product

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
)

func GetProductByID(pid string) (*db.Product, error) {
	product := db.Product{}

	if err := db.GormDB.Where("id = ?", pid).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func FetchProducts(cid string, page int) ([]db.Product, error) {
	products := []db.Product{}

	if err := db.GormDB.
		Limit(10).Offset((page-1)*10).
		Where("category_id = ?", cid).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(productDTO ProductDTO) (*db.Product, error) {
	product := db.Product{
		ID:          utils.GetUUID(),
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
