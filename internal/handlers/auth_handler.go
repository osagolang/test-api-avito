package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/repositories"
	"test-api-avito/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
	userRepo    *repositories.UserRepo
}

func NewAuthHandler(authService *services.AuthService, userRepo *repositories.UserRepo) *AuthHandler {
	return &AuthHandler{authService: authService, userRepo: userRepo}
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

	// Поиск существующего пользователя
	user, err := h.userRepo.FindUser(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
		return
	}

	// Если пользователь не найден, регистрируем
	if user == nil {
		user, err = h.authService.Register(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
			return
		}
		if user == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": "Внутренняя ошибка сервера."})
			return
		}
	}

	// todo тут лишняя проверка для нового пользователя
	token, err := h.authService.Auth(user, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
		return
	}

	// Возвращаем токен в ответ
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
