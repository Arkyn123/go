package service

import (
	"test_service"
	"test_service/pkg/repository"
)

type Users interface {
	CreateUser(user server.User) (server.User, int, error)
	FindAllUsers() ([]server.User, int, error)
	FindById(id int) (server.User, int, error)
}

type Auth interface {
	Login(url string, contentType string, data any) ([]byte, int, error)
	Get() ([]byte, int, error)
}

type Service struct {
	Users
	Auth
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repositories),
		Auth:  NewAuthService(),
	}
}
