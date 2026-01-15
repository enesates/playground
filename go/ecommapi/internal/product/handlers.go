package product

import (
	"fmt"
	"net/http"
	"strconv"

	"ecommapi/internal/auth"
	"ecommapi/internal/helpers/utils"
	"ecommapi/internal/notification"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get products for category
// @Description Get products for a given category starting from a page
// @Tags product
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Success 200 {object} map[string]any "Product Items"
// @Router /products [get]
func GetProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	cid := c.Query("category_id")
	if cid == "" {
		utils.AbortJSON(c, http.StatusBadRequest, "Missing category id")
		return
	}

	products, err := FetchProducts(cid, page)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	items := []gin.H{}

	for _, pr := range products {
		items = append(items, gin.H{
			"name":     pr.Name,
			"price":    pr.Price,
			"category": pr.CategoryID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"total_count": len(items),
	})
}

// AddProduct godoc
// @Summary Create a product
// @Description Create a product for the given category
// @Tags product
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body ProductDTO true "New Product"
// @Success 200 {object} map[string]any "Product details"
// @Router /products [post]
func AddProduct(c *gin.Context) {
	var productDTO ProductDTO

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		utils.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := CreateProduct(productDTO)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	token := c.GetHeader("X-Session-Token")
	session, err := auth.GetSessionByToken(token)
	if err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := notification.CreateNotificationForEvent(session.User.Username, "Product", fmt.Sprintf("Product created: %s", product.ID)); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          product.ID,
		"name":        product.Name,
		"price":       product.Price,
		"description": product.Description,
		"category_id": product.CategoryID,
		"is_active":   product.IsActive,
	})
}
