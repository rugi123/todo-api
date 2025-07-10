package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/todo-api/internal/models"
)

type UserStorage struct {
	pool *pgxpool.Pool
}

func (s UserStorage) Create(ctx context.Context, user models.User) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO users 
    	(username, email, password_hash, created_at, updated_at)
    	VALUES
    	($1, $2, $3, $4, $5)
    	ON CONFLICT (username, email) DO NOTHING`,
		user.UserName, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}
	return nil
}
func (s UserStorage) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET 
			username = $1,
			email = $2,
			password_hash = $3,
			updated_at = $4
		WHERE 
			id = $5
		RETURNING updated_at`
	err := s.pool.QueryRow(ctx, query, user.UserName, user.Email,
		user.PasswordHash, time.Now(), user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}
	return nil
}
func (s UserStorage) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users 
		WHERE id = $1`

	var user models.User
	err := s.pool.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.UserName, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
func (s UserStorage) GetByUserName(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, created_at, updated_at
		FROM users 
		WHERE username = $1`

	var user models.User
	err := s.pool.QueryRow(ctx, query, username).Scan(
		&user.ID, &user.UserName, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}
func (s UserStorage) GetAllTasksForUser(ctx context.Context, user_id int) ([]models.Task, error) {
	query := fmt.Sprintf(`SELECT t.*
		FROM tasks t
		JOIN task_lists tl
		ON t.list_id = tl.id
		WHERE tl.user_id = %v;`, user_id)

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get from db: %w", err)
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err = rows.Scan(&task.ID, &task.ListID, &task.Title, &task.Description, &task.IsComplited,
			&task.DueDate, &task.Priority, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
func (s UserStorage) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}
	return nil
}
