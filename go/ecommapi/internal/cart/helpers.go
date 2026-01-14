package cart

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/product"

	"github.com/gin-gonic/gin"
)

func GenerateProductLineItems(cart *db.Cart, newCartItem *db.CartItem) ([]gin.H, float64, error) {
	pItems := []gin.H{}
	totalAmount := 0.0

	for _, ci := range cart.CartItems {
		product, err := product.GetProductByID(ci.ProductID)

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

func GetOrCreateCart(uid string) (*db.Cart, error) {
	cart, err := FetchCart(uid)

	if err != nil {
		cart, err = CreateCart(uid)

		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}
