package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/todo-api/internal/config"
	"github.com/rugi123/todo-api/internal/models"
)

type PGStorage struct {
	UserStorage     UserStorage
	TaskListStorage TaskListStorage
	TaskStorage     TaskStorage
}

type UserStorage struct {
	pool *pgxpool.Pool
}
type TaskListStorage struct {
	pool *pgxpool.Pool
}
type TaskStorage struct {
	pool *pgxpool.Pool
}

func NewPgStorage(ctx context.Context, cfg *config.PostgresConfig) *PGStorage {
	config, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	//тут теперь можно передлвать pool в функции и работать с таблицами указывая название таблицы query

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
	}
}

// User
func (s UserStorage) Save() error {
	return nil
}
func (s UserStorage) Get() (*models.User, error) {
	return nil, nil
}

// TaskList
func (s TaskListStorage) Save() error {
	return nil
}
func (s TaskListStorage) Get() (*models.TaskList, error) {
	return nil, nil
}

// Task
func (s TaskStorage) Save() error {
	return nil
}
func (s TaskStorage) Get() (*models.Task, error) {
	return nil, nil
}
