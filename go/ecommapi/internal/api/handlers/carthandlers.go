package handlers

import (
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cart, err := dbHelper.GetCart(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	items := []gin.H{}
	totalAmount := 0.0

	for _, ci := range cart.CartItems {
		product, err := dbHelper.GetProductByID(ci.ProductID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		totalAmount += product.Price

		items = append(items, gin.H{
			"product_id": ci.ProductID,
			"quantity":   ci.Quantity,
		})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("X-Session-Token")
	user, err := dbHelper.GetUserByToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cart, err := dbHelper.GetCart(user.ID)
	if err != nil {
		cart, err = dbHelper.CreateCart(user.ID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	cartItem, err := dbHelper.CreateOrUpdateCartItem(cart.ID, cartItemDTO.ProductID, cartItemDTO.Quantity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := []gin.H{
		{
			"product_id": cartItem.ProductID,
			"quantity":   cartItem.Quantity,
		},
	}

	totalAmount := 0.0

	for _, ci := range cart.CartItems {
		product, err := dbHelper.GetProductByID(ci.ProductID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		totalAmount += product.Price

		items = append(items, gin.H{
			"product_id": ci.ProductID,
			"quantity":   ci.Quantity,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"items":        items,
		"total_amount": totalAmount,
	})
}
