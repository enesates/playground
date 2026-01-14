package order

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/product"
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
