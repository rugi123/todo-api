package models

import (
	"time"
)

type Entity interface {
	Validate() error
}

type User struct {
	ID           int       `json:"id"`
	UserName     string    `json:"username"`
	PasswordHash string    `json:"password"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TaskList struct {
	ID          int
	UserID      int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

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

func (u User) Validate() error {
	return nil
}
func (tl TaskList) Validate() error {
	return nil
}
func (t Task) Validate() error {
	return nil
}
