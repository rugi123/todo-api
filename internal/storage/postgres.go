package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/todo-api/internal/config"
)

type PGStorage struct {
	UserStorage     UserStorage
	TaskListStorage TaskListStorage
	TaskStorage     TaskStorage
}

func NewPgStorage(ctx context.Context, cfg *config.PostgresConfig) (*PGStorage, error) {
	config, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	return &PGStorage{
		UserStorage: UserStorage{
			pool: pool,
		},
		TaskListStorage: TaskListStorage{
			pool: pool,
		},
		TaskStorage: TaskStorage{
			pool: pool,
		},
	}, nil
}
