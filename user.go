package server

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
}
