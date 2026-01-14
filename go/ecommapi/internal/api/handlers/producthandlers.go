package handlers

import (
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/dtos"

	"net/http"
	"strconv"

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing category id"})
		return
	}

	products, err := dbHelper.GetProducts(cid, page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product for the given category
// @Tags product
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body dtos.ProductDTO true "New Product"
// @Success 200 {object} map[string]any "Product details"
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var productDTO dtos.ProductDTO

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := dbHelper.CreateProduct(productDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
