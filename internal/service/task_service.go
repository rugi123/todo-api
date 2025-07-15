package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/rugi123/todo-api/internal/models"
	"github.com/rugi123/todo-api/internal/storage"
)

type Service struct {
	Storage storage.PGStorage
}

func NewService(storage storage.PGStorage) *Service {
	return &Service{
		Storage: storage,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}
	return string(bytes), nil
}

func CheckHashPassword(hashed_password string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err
}

func (s Service) Save(ctx context.Context, user *models.User) error {
	if user.ID == 0 {
		password_hash, err := HashPassword(user.PasswordHash)
		if err != nil {
			return fmt.Errorf("failed to generate hash: %w", err)
		}
		user.PasswordHash = password_hash
		return s.Storage.UserStorage.CreateUser(ctx, *user)
	} else {
		return s.Storage.UserStorage.UpdateUser(ctx, user)
	}

}
