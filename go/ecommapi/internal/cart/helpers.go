package cart

import (
	db "ecommapi/internal/helpers/database"
	"math"

	"github.com/gin-gonic/gin"
)

func GenerateProductLineItems(cart *db.Cart) ([]gin.H, float64, error) {
	pItems := []gin.H{}
	totalAmount := 0.0

	for _, ci := range cart.CartItems {
		totalAmount += math.Round((float64(ci.Quantity)*ci.Product.Price)*100) / 100

		pItems = append(pItems, gin.H{
			"product_id": ci.ProductID,
			"quantity":   ci.Quantity,
		})
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
