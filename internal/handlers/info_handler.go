package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/services"
)

type InfoHandler struct {
	infoService *services.InfoService
}

func NewInfoHandler(infoService *services.InfoService) *InfoHandler {
	return &InfoHandler{infoService: infoService}
}

func (h *InfoHandler) GetUserInfo(c *gin.Context) {

	// Идём в контекст (в Middleware) за ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
		return
	}

	// Получаем количество монет
	coins, err := h.infoService.UserInfo(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"coins": coins,
	})
}
