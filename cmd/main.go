package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rugi123/todo-api/internal/config"
	"github.com/rugi123/todo-api/internal/handlers"
	"github.com/rugi123/todo-api/internal/service"
	"github.com/rugi123/todo-api/internal/storage"
)

func main() {
	cfg, err := config.Load("internal/config/config.yaml")
	if err != nil {
		fmt.Println("failed to load: ", err)
	}

	PGStorage, err := storage.NewPgStorage(context.Background(), &cfg.PostgresConfig)
	if err != nil {
		fmt.Println("failed to create pg storage: ", err)
	}
	service := service.NewService(*PGStorage)

	authHandler := handlers.NewAuthHandler(*cfg, *service)

	router := gin.Default()

	router.LoadHTMLGlob("static/templates/*.html")
	router.Static("/static", "./static")

	authHandler.RegisterRoutes(router)

	router.Run(":" + cfg.AppConfig.Port)
}
