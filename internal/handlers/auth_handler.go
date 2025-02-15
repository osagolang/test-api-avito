package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-api-avito/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login обрабатывает запрос авторизации/регистрации пользователя
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный запрос"})
		return
	}

	user, token, err := h.authService.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидные данные"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
