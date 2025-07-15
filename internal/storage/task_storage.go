package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/todo-api/internal/models"
)

type TaskStorage struct {
	pool *pgxpool.Pool
}

func (s TaskStorage) CreateTask(ctx context.Context, task models.Task) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO tasks 
		(list_id, title, description, is_completed, due_date, priority, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (list_id) DO NOTHING`,
		task.ListID, task.Title, task.Description, task.IsComplited,
		task.DueDate, task.Priority, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}
func (s TaskStorage) UpdateTask(ctx context.Context, task *models.Task) error {
	query := `
		UPDATE tasks 
		SET 
			title = $1,
			description = $2,
			is_completed = $3,
			due_date = $4
			priority = $5
			updated_at = $6
		WHERE 
			id = $7
		RETURNING updated_at`
	err := s.pool.QueryRow(ctx, query, task.Title, task.Description,
		task.IsComplited, task.DueDate, task.Priority, time.Now(), task.ID).Scan(&task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}
	return nil
}
func (s TaskStorage) GetByID(ctx context.Context, id int) (*models.Task, error) {
	query := `
		SELECT id, list_id ,title, description, is_completed, due_date, priority, created_at, updated_at
		FROM tasks 
		WHERE id = $1`

	var task models.Task
	err := s.pool.QueryRow(ctx, query, task.ID).Scan(&task.ID, &task.ListID, &task.Title, &task.Description,
		&task.IsComplited, &task.DueDate, &task.Priority, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}
func (s TaskStorage) GetByListID(ctx context.Context, list_id int) (*models.Task, error) {
	query := `
		SELECT id, list_id ,title, description, is_completed, due_date, priority, created_at, updated_at
		FROM tasks 
		WHERE list_id = $1`

	var task models.Task
	err := s.pool.QueryRow(ctx, query, list_id).Scan(&task.ID, &task.ListID, &task.Title, &task.Description,
		&task.IsComplited, &task.DueDate, &task.Priority, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}
func (s TaskStorage) DeleteTask(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}
	return nil
}
