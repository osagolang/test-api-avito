package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/services"
	"test-api-avito/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Auth обрабатывает запрос авторизации/регистрации пользователя
func (h *AuthHandler) Auth(c *gin.Context) {
	// Парсим тело запроса
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Неверный запрос."})
		return
	}

	// Пробуем авторизовать
	user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		// Регистрируем, если не найден
		user, err = h.authService.Register(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
			return
		}
	}

	// Генерируем токен
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
		return
	}

	// Возвращаем токен в ответ
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
