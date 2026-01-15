package order

import (
	"ecommapi/internal/auth"
	"ecommapi/internal/cart"
	"ecommapi/internal/helpers/utils"
	notif "ecommapi/internal/notification"

	"net/http"

	"github.com/gin-gonic/gin"
)

// PlaceOrder godoc
// @Summary Create order
// @Description Create order from the cart
// @Tags order
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body OrderDTO true "New Order Item"
// @Success 200 {object} map[string]any "Order items"
// @Router /orders [post]
func PlaceOrder(c *gin.Context) {
	var orderDTO OrderDTO

	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		utils.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if !HasEnoughStock(orderDTO.Items) {
		utils.AbortJSON(c, http.StatusBadRequest, "Not enough stock")
		return
	}

	token := c.GetHeader("X-Session-Token")
	session, err := auth.GetSessionByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err := CreateOrder(orderDTO, session.User.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := cart.DeleteCartByUserID(session.User.ID); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := UpdateProductInventories(orderDTO.Items); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notif.CreateEventNotif(session.User.ID, "Cart Update", "Cart is removed after order is placed"); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notif.CreateEventNotif(session.User.ID, "Inventory Update", "Stocks are updated after order is placed"); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err = FetchOrder(order.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, totalAmount, err := GenerateOrderLineItems(order)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := UpdateOrderTotal(order.ID, totalAmount); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"items":            items,
		"shipping_street":  orderDTO.ShippingStreet,
		"shipping_city":    orderDTO.ShippingCity,
		"shipping_zip":     orderDTO.ShippingZip,
		"shipping_country": orderDTO.ShippingCountry,
		"total_amount":     totalAmount,
	})
}
