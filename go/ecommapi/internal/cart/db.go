package cart

import (
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/helpers/utils"
	"math"

	"gorm.io/gorm"
)

func FetchCart(uid string) (*db.Cart, error) {
	cart := db.Cart{}

	if err := db.GormDB.
		Preload("CartItems").
		Preload("CartItems.Product").
		Where("user_id = ?", uid).
		First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func CreateCart(uid string) (*db.Cart, error) {
	cart := db.Cart{
		ID:          utils.GetUUID(),
		UserID:      uid,
		TotalAmount: 0.0,
	}

	if err := db.GormDB.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func GetCartItem(cid string, pid string) (*db.CartItem, error) {
	cartItem := db.CartItem{}

	if err := db.GormDB.
		Preload("Product").
		Where("cart_id = ? AND product_id = ?", cid, pid).
		First(&cartItem).Error; err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func CreateOrUpdateCartItem(cid string, pid string, quantity int) (*db.CartItem, error) {
	cartItem, _ := GetCartItem(cid, pid)

	if cartItem == nil {
		cartItem = &db.CartItem{
			ID:        utils.GetUUID(),
			CartID:    cid,
			ProductID: pid,
			Quantity:  quantity,
		}

		if err := db.GormDB.Create(&cartItem).Error; err != nil {
			return nil, err
		}

		cartItem, _ = GetCartItem(cid, pid)
	} else if err := db.GormDB.
		Model(&cartItem).
		Where("id = ?", cartItem.ID).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity)).
		Error; err != nil {
		return nil, err
	}

	totalAmount := math.Round((float64(quantity)*cartItem.Product.Price)*100) / 100

	if err := UpdateCartTotal(cid, totalAmount); err != nil {
		return nil, err
	}

	return cartItem, nil
}

func DeleteCartByUserID(uid string) error {
	if err := db.GormDB.Unscoped().Where("user_id = ?", uid).Delete(&db.Cart{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateCartTotal(cid string, price float64) error {
	if err := db.GormDB.
		Model(&db.Cart{}).
		Where("id = ?", cid).
		UpdateColumn("total_amount", gorm.Expr("total_amount + ?", price)).
		Error; err != nil {
		return err
	}

	return nil
}
