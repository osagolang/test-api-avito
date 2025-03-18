package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// Брать секретный ключ надо из фала конфигурации os.Getenv
// Но я для простоты и удобства проверки написал простую строку (на боевом проекте так делать нельзя!)
var jwtSecret = []byte("avito-secret")

func GenerateToken(userID int) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return jwtSecret, nil
	})

	return token, err
}
