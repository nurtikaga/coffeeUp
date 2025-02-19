package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Sslmode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode))
	if err != nil {
		fmt.Println("1st")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("2nd")
		return nil, err
	}

	return db, nil
}
