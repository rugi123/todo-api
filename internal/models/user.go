package models

import "time"

type User struct {
	ID           int
	UserName     string
	PasswordHash string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
