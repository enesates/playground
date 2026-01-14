package handlers

import (
	h "ecommapi/internal/api/helpers"
	dbHelper "ecommapi/internal/database/helpers"
	"ecommapi/internal/dtos"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInventory godoc
// @Summary Get product inventory
// @Description Get inventory details of a product
// @Tags inventory
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Success 200 {object} map[string]any "Inventory details"
// @Router /inventory/:product_id [get]
func GetInventory(c *gin.Context) {
	pid := c.Param("product_id")

	inventory, err := dbHelper.GetInventory(pid)
	if inventory == nil || err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quantity":        inventory.Quantity,
		"stock":           inventory.Quantity - inventory.Reserved,
		"storageLocation": inventory.Location,
	})
}

// UpdateInventory godoc
// @Summary Update inventory
// @Description Increase product inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body dtos.StockDTO true "Stock Data"
// @Success 200 {object} map[string]any "Inventory details"
// @Router /inventory/:product_id [post]
func UpdateInventory(c *gin.Context) {
	var stockDTO dtos.StockDTO
	var err error
	pid := c.Param("product_id")

	if err := c.ShouldBindJSON(&stockDTO); err != nil {
		h.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if stockDTO.IncerementBy < 0 {
		h.AbortJSON(c, http.StatusBadRequest, "Invalid quantity")
		return
	}

	inventory, _ := dbHelper.GetInventory(pid)
	if inventory == nil {
		inventory, err = dbHelper.CreateInventory(pid, stockDTO)
	} else {
		inventory, err = dbHelper.UpdateInventory(pid, stockDTO)
	}

	if err != nil {
		h.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quantity":        inventory.Quantity,
		"stock":           inventory.Quantity - inventory.Reserved,
		"storageLocation": inventory.Location,
	})
}
