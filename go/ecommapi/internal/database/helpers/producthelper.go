package db

import (
	db "ecommapi/internal/database"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"

	"errors"
	"log"
)

func GetProducts(categoryId string, page int) ([]models.Product, error) {
	products := []models.Product{}
	productsResult := db.GormDB.Limit(10).Offset((page-1)*10).Where("category_id = ?", categoryId).Find(&products)

	if productsResult.Error != nil {
		log.Printf("Error getting products: %v", productsResult.Error)
		return products, errors.New("Internal Server Error")
	}

	return products, nil

}

func CreateProduct(productDTO models.ProductDTO) (*models.Product, error) {
	product := models.Product{
		ID:          helpers.GetUUID(),
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		CategoryID:  productDTO.CategoryID,
		IsActive:    true,
	}

	productResult := db.GormDB.Create(&product)

	if productResult.Error != nil {
		log.Printf("Error creating product: %v", productResult.Error)
		return nil, errors.New("Internal Server Error")
	}

	log.Printf("Product created successfully with ID: %s", product.ID)
	return &product, nil

}

//   ```json
//   ```
// - **Response Example (201 Created):**
//   ```json
//   {
//     "items": [
//       {
//         "name": "string",
//         "price": decimal,
//         "category": "string"
//       },
//       ...
//     ],
//     "total_count": decimal
//   }
