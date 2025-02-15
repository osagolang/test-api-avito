package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"test-api-avito/internal/handlers"
	"test-api-avito/internal/repositories"
	"test-api-avito/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=db-avito user=dev password=dev sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка соединения с БД: %v", err)
	}

	fmt.Println("Успешное подключение к БД")

	userRepo := repositories.NewUserRepo(db)
	walletRepo := repositories.NewWalletRepo(db)

	authService := services.NewAuthService(userRepo, walletRepo)

	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	authGroup := router.Group("/api")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	fmt.Println("Сервер: http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Ошибка запуска сервера: %v", err)
	}
}
