package product

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductByID_Success(t *testing.T) {
	db.SetupTestDB()

	pid := utils.GetUUID()
	product := db.Product{
		ID:          pid,
		Name:        "product-name",
		Description: "product-description",
		Price:       12.34,
		CategoryID:  "product-category",
	}

	db.GormDB.Create(&product)
	p, err := GetProductByID(pid)

	assert.NoError(t, err)
	assert.Equal(t, product.Name, p.Name)
}

func TestGetProductByID_NotFound(t *testing.T) {
	db.SetupTestDB()

	product := db.Product{
		ID:          "pid",
		Name:        "product-name",
		Description: "product-description",
		Price:       12.34,
		CategoryID:  "product-category",
	}

	db.GormDB.Create(&product)
	p, err := GetProductByID("wrong-id")

	assert.Error(t, err)
	assert.Nil(t, p)
}

func TestFetchProducts(t *testing.T) {
	db.SetupTestDB()

	pid1 := utils.GetUUID()
	product1 := db.Product{
		ID:          pid1,
		Name:        "product-name",
		Description: "product-description",
		Price:       12.34,
		CategoryID:  "product-category",
	}
	db.GormDB.Create(&product1)

	pid2 := utils.GetUUID()
	product2 := db.Product{
		ID:          pid2,
		Name:        "product-name",
		Description: "product-description",
		Price:       12.34,
		CategoryID:  "product-category",
	}
	db.GormDB.Create(&product2)

	products, err := FetchProducts("product-category", 1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(products))
}

func TestCreateProduct(t *testing.T) {
	db.SetupTestDB()

	product := ProductDTO{
		Name:        "product-name",
		Description: "product-description",
		Price:       12.34,
		CategoryID:  "product-category",
	}

	p, err := CreateProduct(product)

	assert.NoError(t, err)
	assert.Equal(t, product.Name, p.Name)
}
