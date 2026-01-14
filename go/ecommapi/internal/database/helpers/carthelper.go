package db

import (
	db "ecommapi/internal/database"
	"ecommapi/internal/helpers"
	"ecommapi/internal/models"

	"gorm.io/gorm"
)

func GetCart(uid string) (*models.Cart, error) {
	cart := models.Cart{}

	if err := db.GormDB.
		Where("user_id = ?", uid).
		First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func CreateCart(uid string) (*models.Cart, error) {
	cart := models.Cart{
		ID:          helpers.GetUUID(),
		UserID:      uid,
		TotalAmount: 0.0,
	}

	if err := db.GormDB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func GetCartItem(cid string, pid string) (*models.CartItem, error) {
	cartItem := models.CartItem{}

	if err := db.GormDB.
		Where("cart_id = ? AND product_id = ?", cid, pid).
		First(&cartItem).Error; err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func CreateOrUpdateCartItem(cid string, pid string, quantity int) (*models.CartItem, error) {
	cartItem, _ := GetCartItem(cid, pid)

	if cartItem == nil {
		cartItem = &models.CartItem{
			ID:        helpers.GetUUID(),
			CartID:    cid,
			ProductID: pid,
			Quantity:  quantity,
		}

		if err := db.GormDB.Create(&cartItem).Error; err != nil {
			return nil, err
		}

		return cartItem, nil
	}

	if err := db.GormDB.
		Model(&cartItem).
		Where("id = ?", cartItem.ID).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", cartItem.Quantity+quantity)).
		Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}
