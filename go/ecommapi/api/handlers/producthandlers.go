package handlers

import (
	dbHelper "ecommapi/internal/database/helpers"
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

	c.JSON(http.StatusOK, gin.H{
		// "items": [
		//   gin.H{
		//     "name": "string",
		//     "price": decimal,
		//     "category": "string",
		//   },
		//   ///...
		// ],
		"a":           products[0],
		"total_count": "decimal",
	})
}
