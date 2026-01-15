package inventory

import (
	"ecommapi/internal/auth"
	"ecommapi/internal/helpers/utils"
	notif "ecommapi/internal/notification"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInventory godoc
// @Summary Get product inventory
// @Description Get inventory details of a product
// @Tags inventory
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param product_id path string true "Product ID"
// @Success 200 {object} map[string]any "Inventory details"
// @Router /inventory/:product_id [get]
func GetInventory(c *gin.Context) {
	pid := c.Param("product_id")

	inventory, err := FetchInventory(pid)
	if inventory == nil || err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quantity":        inventory.Quantity,
		"stock":           inventory.Quantity - inventory.Reserved,
		"storageLocation": inventory.Location,
	})
}

// CreateOrUpdateInventory godoc
// @Summary Update inventory
// @Description Increase product inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param X-Session-Token header string true "Session token"
// @Param data body StockDTO true "Stock Data"
// @Param product_id path string true "Product ID"
// @Success 200 {object} map[string]any "Inventory details"
// @Router /inventory/:product_id [patch]
func CreateOrUpdateInventory(c *gin.Context) {
	var stockDTO StockDTO
	var err error
	pid := c.Param("product_id")

	if err := c.ShouldBindJSON(&stockDTO); err != nil {
		utils.AbortJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if stockDTO.IncerementBy < 0 {
		utils.AbortJSON(c, http.StatusBadRequest, "Invalid quantity")
		return
	}

	inventory, _ := FetchInventory(pid)
	if inventory == nil {
		inventory, err = CreateInventory(pid, stockDTO)
	} else {
		inventory, err = UpdateInventory(pid, stockDTO)
	}

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

	if err := notif.CreateEventNotif(session.User.ID, "Inventory Update", fmt.Sprintf("Stock is increased by %d for %s", stockDTO.IncerementBy, pid)); err != nil {
		utils.AbortJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quantity":        inventory.Quantity,
		"stock":           inventory.Quantity - inventory.Reserved,
		"storageLocation": inventory.Location,
	})
}
