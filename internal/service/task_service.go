package service

import (
	"github.com/rugi123/todo-api/internal/models"
	"github.com/rugi123/todo-api/internal/storage"
)

type Storage[T models.Entity] interface {
	Save() error
	Get() (*T, error)
}

type Service struct {
	UserService     UserService
	TaskListService TaskListService
	TaskService     TaskService
}

type UserService struct {
	storage Storage[models.User]
}
type TaskListService struct {
	storage Storage[models.TaskList]
}
type TaskService struct {
	storage Storage[models.Task]
}

func NewService[T models.Entity](storage storage.PGStorage) *Service {
	return &Service{
		UserService: UserService{
			storage: storage.UserStorage,
		},
		TaskListService: TaskListService{
			storage: storage.TaskListStorage,
		},
		TaskService: TaskService{
			storage: storage.TaskStorage,
		},
	}
}

func (s *UserService) GenHashPasswd() {
	s.storage.Save()
}
