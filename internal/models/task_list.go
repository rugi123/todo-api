package models

import "time"

type TaskList struct {
	ID          int
	UserID      int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
