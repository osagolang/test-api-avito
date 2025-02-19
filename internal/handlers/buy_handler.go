package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/repositories"
)

type BuyHandler struct {
	buyRepo *repositories.BuyRepo
}

func NewBuyHandler(buyRepo *repositories.BuyRepo) *BuyHandler {
	return &BuyHandler{buyRepo: buyRepo}
}

func (h *BuyHandler) BuyItem(c *gin.Context) {
	userID, _ := c.Get("userID")
	itemName := c.Param("item")

	err := h.buyRepo.BuyItem(userID.(int), itemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Покупка успешна"})
}
