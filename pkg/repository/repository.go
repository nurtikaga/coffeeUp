package repository

import (
	"github.com/jmoiron/sqlx"
	coffeeup "github.com/nurtikaga/coffeeUp"
)

const (
	userTable = "users"
)

type Authorization interface {
	CreateUser(user coffeeup.User) (int, error)
	GetUser(username, password string) (coffeeup.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
	}
}
