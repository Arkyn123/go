package repository

import (
	"github.com/jmoiron/sqlx"
	"net/http"
	"test_service"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user server.User) (server.User, int, error) {
	query := "INSERT INTO users (name) VALUES ($1) RETURNING *"

	var response server.User
	if err := r.db.QueryRow(query, user.Name).Scan(&response.Id, &response.Name); err != nil {
		return server.User{}, http.StatusBadRequest, err
	}

	return response, http.StatusCreated, nil
}

func (r *UserPostgres) FindByName(name string) (server.User, int, error) {
	query := "SELECT * from users WHERE name = $1"

	var user server.User
	if err := r.db.QueryRow(query, name).Scan(&user.Id, &user.Name); err != nil {
		return server.User{}, http.StatusBadRequest, nil
	}

	return user, http.StatusOK, nil
}

func (r *UserPostgres) FindAllUsers() ([]server.User, int, error) {
	query := "SELECT * FROM users"
	var users []server.User
	rows, err := r.db.Query(query)
	if err != nil {
		return users, http.StatusBadRequest, err
	}
	defer rows.Close()

	for rows.Next() {
		var user server.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return users, http.StatusBadRequest, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return users, http.StatusBadRequest, err
	}

	return users, http.StatusOK, nil
}

func (r *UserPostgres) FindById(id int) (server.User, int, error) {
	query := "SELECT * from users WHERE id = $1"
	var user server.User
	if err := r.db.Get(&user, query, id); err != nil {
		return user, http.StatusNotFound, err
	}

	return user, http.StatusOK, nil
}
