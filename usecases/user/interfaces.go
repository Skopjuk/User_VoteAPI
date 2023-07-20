package user

import (
	"userapi/models"
)

type InsertUser interface {
	InsertUser(user *models.User) error
}

type FindUserByUsername interface {
	FindUserByUsername(username string) (*models.User, error)
}

type DeleteUser interface {
	DeleteUser(username string) error
}

type UpdateUser interface {
	UpdateUser(user *models.User) error
}

type AuthenticateUser interface {
	AuthenticateUser(username string, password string) bool
}

type ChangeUsersPassword interface {
	ChangeUsersPassword(username string, password string) bool
}
