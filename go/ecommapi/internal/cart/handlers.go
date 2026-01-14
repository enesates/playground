package cart

import (
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/user"

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
	user, err := user.GetUserByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cart, err := FetchCart(user.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusNotFound, err.Error())
		return
	}

	items, totalAmount, err := GenerateProductLineItems(cart, nil)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
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
// @Param data body CartItemDTO true "New Cart Item"
// @Success 200 {object} map[string]any "Cart items"
// @Router /cart/items [post]
func AddToCart(c *gin.Context) {
	var cartItemDTO CartItemDTO

	if err := c.ShouldBindJSON(&cartItemDTO); err != nil {
		utils.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	token := c.GetHeader("X-Session-Token")
	user, err := user.GetUserByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cart, err := GetOrCreateCart(user.ID)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	cartItem, err := CreateOrUpdateCartItem(cart.ID, cartItemDTO.ProductID, cartItemDTO.Quantity)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, totalAmount, err := GenerateProductLineItems(cart, cartItem)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"items":        items,
		"total_amount": totalAmount,
	})
}
