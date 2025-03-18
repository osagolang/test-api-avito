package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/repositories"
)

type TransferHandler struct {
	transferCoinsRepo *repositories.TransferCoinsRepo
	userRepo          *repositories.UserRepo
}

func NewTransferHandler(transferCoinsRepo *repositories.TransferCoinsRepo, userRepo *repositories.UserRepo) *TransferHandler {
	return &TransferHandler{transferCoinsRepo: transferCoinsRepo, userRepo: userRepo}
}

func (h *TransferHandler) SendCoin(c *gin.Context) {

	fromUserID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
		return
	}

	type Request struct {
		Username string `json:"username" binding:"required"`
		Amount   int    `json:"amount" binding:"required"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Неверный запрос."})
		return
	}

	toUserID, err := h.userRepo.FindUser(req.Username)

	err = h.transferCoinsRepo.TransferCoins(fromUserID.(int), toUserID.ID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Успешный перевод."})
}
