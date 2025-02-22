package service

import (
	coffeeup "github.com/nurtikaga/coffeeUp"
	"github.com/nurtikaga/coffeeUp/pkg/repository"
)

type Authorization interface {
	CreateUser(user coffeeup.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
