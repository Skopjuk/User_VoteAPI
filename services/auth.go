package services

import "userapi/repository"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}
