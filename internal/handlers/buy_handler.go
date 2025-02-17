package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/services"
)

type BuyHandler struct {
	buyService *services.BuyService
}

func NewBuyHandler(buyService *services.BuyService) *BuyHandler {
	return &BuyHandler{buyService: buyService}
}

func (h *BuyHandler) BuyItem(c *gin.Context) {
	userID, _ := c.Get("userID")
	itemName := c.Param("item")

	err := h.buyService.BuyItem(userID.(int), itemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Покупка успешна"})
}
