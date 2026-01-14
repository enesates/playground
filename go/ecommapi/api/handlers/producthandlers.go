package handlers

import (
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/models"

	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	categoryId := c.Query("category_id")
	if categoryId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(errors.New("Missing category id")))
		return
	}

	products, err := dbHelper.GetProducts(categoryId, page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
		return
	}

	type Item struct {
		name     string
		price    float64
		category string
	}

	items := []Item{}

	for _, pr := range products {
		item := Item{
			name:     pr.Name,
			price:    pr.Price,
			category: pr.CategoryID,
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"total_count": len(items),
	})
}

func CreateProduct(c *gin.Context) {
	var productDTO models.ProductDTO

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, c.Error(err))
		return
	}

	product, err := dbHelper.CreateProduct(productDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
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
