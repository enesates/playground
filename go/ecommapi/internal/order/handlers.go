package order

import (
	"ecommapi/internal/cart"
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/user"

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
	userObj, err := user.GetUserByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err := CreateOrder(orderDTO, userObj.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := cart.DeleteCartByUserID(userObj.ID); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := UpdateProductInventories(orderDTO.Items); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, totalAmount, err := GenerateOrderLineItems(order)
	if err != nil {
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
