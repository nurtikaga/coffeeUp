package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	coffeeup "github.com/nurtikaga/coffeeUp"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user coffeeup.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, number, password_hash, role) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := s.db.QueryRow(query, user.Name, user.Email, user.Number, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (coffeeup.User, error) {
	var user coffeeup.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
