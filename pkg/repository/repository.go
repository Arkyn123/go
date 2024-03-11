package repository

import (
	"github.com/jmoiron/sqlx"
	"test_service"
)

type Users interface {
	CreateUser(user server.User) (server.User, int, error)
	FindByName(name string) (server.User, int, error)
	FindAllUsers() ([]server.User, int, error)
	FindById(id int) (server.User, int, error)
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUserPostgres(db),
	}
}
