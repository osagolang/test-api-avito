package main

import (
	"database/sql"
	"fmt"
	"log"
	"test-api-avito/internal/handlers"
	"test-api-avito/internal/middleware"
	"test-api-avito/internal/repositories"
	"test-api-avito/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=db-avito user=dev password=dev sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	fmt.Println("Успешное подключение к БД")

	userRepo := repositories.NewUserRepo(db)
	userInfoRepo := repositories.NewUserInfoRepo(db)
	itemRepo := repositories.NewItemRepo(db)

	authService := services.NewAuthService(userRepo)
	infoService := services.NewInfoService(userInfoRepo)
	buyService := services.NewBuyService(userInfoRepo, itemRepo)

	authHandler := handlers.NewAuthHandler(authService, userRepo)
	infoHandler := handlers.NewInfoHandler(infoService)
	buyHandler := handlers.NewBuyHandler(buyService)

	router := gin.Default()

	router.POST("/api/auth", authHandler.Auth)
	router.GET("/api/info", middleware.AuthMiddleware(), infoHandler.GetUserInfo)
	router.GET("api/buy/:item", middleware.AuthMiddleware(), buyHandler.BuyItem)

	fmt.Println("Сервер: http://localhost:8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
