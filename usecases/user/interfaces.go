package user

import (
	"userapi/models"
)

type InsertUser interface {
	InsertUser(user models.User) (id int, err error)
}

type FindUserByUsername interface {
	FindUserByUsername(username string) (models.User, error)
}

type DeleteUser interface {
	DeleteUser(username string) error
}

type UpdateUser interface {
	UpdateUser(user models.User, id int) error
}

type AuthenticateUser interface {
	AuthenticateUser(username string, password string) bool
}

type ChangeUsersPassword interface {
	ChangeUsersPassword(username string, password string) bool
}

type GetAll interface {
	GetAll() (usersList []models.User, err error)
}

type GetUserById interface {
	GetUserById(id int) (user models.User, err error)
}
