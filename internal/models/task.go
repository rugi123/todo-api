package models

import "time"

type Task struct {
	ID          int
	ListID      int
	Title       string
	Description string
	IsComplited bool
	DueDate     time.Time
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
