package main

import (
	"context"
	"fmt"

	"github.com/rugi123/todo-api/internal/config"
	"github.com/rugi123/todo-api/internal/storage"
)

func main() {
	cfg, err := config.Load("internal/config/config.yaml")
	if err != nil {
		fmt.Println("failed to load: ", err)
	}

	storage, err := storage.NewPgStorage(context.Background(), &cfg.PostgresConfig)
	if err != nil {
		fmt.Println("failed to create pg storage: ", err)
	}
	//service := service.NewService(*storage)
}
