package order

import (
	"ecommapi/internal/cart"
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/inventory"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateOrderLineItems(order *db.Order) ([]gin.H, float64, error) {
	oItems := []gin.H{}
	totalAmount := 0.0

	for _, oi := range order.OrderItems {
		totalAmount += math.Round((float64(oi.Quantity)*oi.UnitPrice)*100) / 100

		oItems = append(oItems, gin.H{
			"product_id": oi.ProductID,
			"quantity":   oi.Quantity,
		})
	}

	return oItems, totalAmount, nil
}

func HasEnoughStock(items []cart.CartItemDTO) bool {
	for _, i := range items {
		stock, err := inventory.FetchInventory(i.ProductID)

		if err != nil {
			return false
		}

		if stock.Quantity-(stock.Reserved+i.Quantity) < 0 {
			return false
		}
	}

	return true
}

func UpdateProductInventories(items []cart.CartItemDTO) error {
	for _, i := range items {
		stock, err := inventory.FetchInventory(i.ProductID)

		if err != nil {
			return err
		}

		if err = db.GormDB.
			Model(&stock).
			Where("id = ?", stock.ID).
			UpdateColumn("reserved", gorm.Expr("reserved + ?", i.Quantity)).
			Error; err != nil {
			return err
		}
	}

	return nil
}
