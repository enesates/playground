package helpers

import (
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/models"

	"github.com/gin-gonic/gin"
)

func GenerateProductLineItems(cart *models.Cart, newCartItem *models.CartItem) ([]gin.H, float64, error) {
	pItems := []gin.H{}
	totalAmount := 0.0

	for _, ci := range cart.CartItems {
		product, err := dbHelper.GetProductByID(ci.ProductID)

		if err != nil {
			return nil, 0.0, err
		}

		totalAmount += product.Price

		pItems = append(pItems, gin.H{
			"product_id": ci.ProductID,
			"quantity":   ci.Quantity,
		})
	}

	if newCartItem != nil {
		pItems = append(pItems, gin.H{
			"product_id": newCartItem.ProductID,
			"quantity":   newCartItem.Quantity,
		})
		totalAmount += newCartItem.Product.Price
	}

	return pItems, totalAmount, nil
}

func GetOrCreateCart(uid string) (*models.Cart, error) {
	cart, err := dbHelper.GetCart(uid)

	if err != nil {
		cart, err = dbHelper.CreateCart(uid)

		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}
