package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"test-api-avito/internal/utils"
)

// AuthMiddleware - проверка JWT-токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовка
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
			c.Abort()
			return
		}

		// Парсим токен
		token, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
			c.Abort()
			return
		}

		// Извлекаем ID пользователя из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
			c.Abort()
			return
		}

		userID, ok := claims["sub"].(float64) // JWT возвращает float64 для чисел
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован."})
			c.Abort()
			return
		}

		// Добавляем ID пользователя в контекст
		c.Set("userID", int(userID))
		c.Next()
	}
}
