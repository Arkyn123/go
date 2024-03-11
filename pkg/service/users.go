package service

import (
	"errors"
	"net/http"
	"test_service"
	"test_service/pkg/repository"
)

type UsersService struct {
	repository repository.Users
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{repository: repository}
}

func (s *UsersService) CreateUser(user server.User) (server.User, int, error) {
	existingUser, statusCode, err := s.repository.FindByName(user.Name)
	if err != nil {
		return server.User{}, statusCode, err
	}

	if existingUser != (server.User{}) {
		return server.User{}, http.StatusConflict, errors.New("user already exists")
	}

	return s.repository.CreateUser(user)
}

func (s *UsersService) FindAllUsers() ([]server.User, int, error) {
	return s.repository.FindAllUsers()
}

func (s *UsersService) FindById(id int) (server.User, int, error) {
	return s.repository.FindById(id)
}
