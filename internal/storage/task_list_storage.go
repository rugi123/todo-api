package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/todo-api/internal/models"
)

type TaskListStorage struct {
	pool *pgxpool.Pool
}

func (s TaskListStorage) Create(ctx context.Context, taskList models.TaskList) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO task_lists 
		(user_id, title, description, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO NOTHING`,
		taskList.UserID, taskList.Title, taskList.Description, taskList.CreatedAt, taskList.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}
func (s TaskListStorage) Update(ctx context.Context, task_lists *models.TaskList) error {
	query := `
		UPDATE task_lists 
		SET 
			title = $1,
			description = $2,
			updated_at = $3
		WHERE 
			id = $4
		RETURNING updated_at`
	err := s.pool.QueryRow(ctx, query, task_lists.Title, task_lists.Description,
		time.Now(), task_lists.ID).Scan(&task_lists.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}
	return nil
}
func (s TaskListStorage) GetByID(ctx context.Context, id int) (*models.TaskList, error) {
	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM task_lists 
		WHERE id = $1`

	var task_list models.TaskList
	err := s.pool.QueryRow(ctx, query, id).Scan(&task_list.ID, &task_list.UserID,
		&task_list.Title, &task_list.Description, &task_list.CreatedAt, &task_list.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &task_list, nil
}
func (s TaskListStorage) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM task_lists WHERE id = $1"
	_, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}
	return nil
}
