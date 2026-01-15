package order

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/product"

	"gorm.io/gorm"
)

func CreateOrder(orderDTO OrderDTO, uid string) (*db.Order, error) {
	order := db.Order{
		ID:              utils.GetUUID(),
		UserID:          uid,
		Status:          "done",
		TotalAmount:     0.0,
		ShippingStreet:  orderDTO.ShippingStreet,
		ShippingCity:    orderDTO.ShippingCity,
		ShippingZip:     orderDTO.ShippingZip,
		ShippingCountry: orderDTO.ShippingCountry,
	}

	if err := db.GormDB.Create(&order).Error; err != nil {
		return nil, err
	}

	for _, oi := range orderDTO.Items {
		product, err := product.GetProductByID(oi.ProductID)

		if err != nil {
			return nil, err
		}

		orderItem := db.OrderItem{
			ID:        utils.GetUUID(),
			OrderID:   order.ID,
			ProductID: oi.ProductID,
			Quantity:  oi.Quantity,
			UnitPrice: product.Price,
		}

		if err := db.GormDB.Create(&orderItem).Error; err != nil {
			return nil, err
		}
	}

	return &order, nil
}

func FetchOrder(uid string) (*db.Order, error) {
	order := db.Order{}

	if err := db.GormDB.
		Preload("OrderItems").
		Where("id = ?", uid).
		First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func UpdateOrderTotal(oid string, price float64) error {
	if err := db.GormDB.
		Model(&db.Order{}).
		Where("id = ?", oid).
		UpdateColumn("total_amount", gorm.Expr("total_amount + ?", price)).
		Error; err != nil {
		return err
	}

	return nil
}
