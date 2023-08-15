package services

import (
	"userapi/models"
	"userapi/repository"
)

type Authorisation interface {
	CreateUser(user models.User)
}

type Service struct {
	Authorisation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorisation: NewAuthService(repos.Authorisation)}
}
