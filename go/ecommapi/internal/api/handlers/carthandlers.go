package handlers

import (
	"ecommapi/internal/api/helpers"
	h "ecommapi/internal/api/helpers"
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/dtos"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCart godoc
// @Summary Get cart for customer
// @Description Get cart items for customer
// @Tags cart
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Success 200 {object} map[string]any "Cart Items"
// @Router /cart [get]
func GetCart(c *gin.Context) {
	token := c.GetHeader("X-Session-Token")
	user, err := dbHelper.GetUserByToken(token)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cart, err := dbHelper.GetCart(user.ID)
	if err != nil {
		h.AbortJSON(c, http.StatusNotFound, err.Error())
		return
	}

	items, totalAmount, err := helpers.GenerateProductLineItems(cart, nil)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"items":        items,
		"total_amount": totalAmount,
	})
}

// AddToCart godoc
// @Summary Add to cart
// @Description Add product to the cart
// @Tags cart
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body dtos.CartItemDTO true "New Cart Item"
// @Success 200 {object} map[string]any "Cart items"
// @Router /cart/items [post]
func AddToCart(c *gin.Context) {
	var cartItemDTO dtos.CartItemDTO

	if err := c.ShouldBindJSON(&cartItemDTO); err != nil {
		h.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	token := c.GetHeader("X-Session-Token")
	user, err := dbHelper.GetUserByToken(token)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cart, err := helpers.GetOrCreateCart(user.ID)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cartItem, err := dbHelper.CreateOrUpdateCartItem(cart.ID, cartItemDTO.ProductID, cartItemDTO.Quantity)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, totalAmount, err := helpers.GenerateProductLineItems(cart, cartItem)
	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"items":        items,
		"total_amount": totalAmount,
	})
}
