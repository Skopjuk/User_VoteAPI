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
	ChangeUsersPassword(id int, password string) error
}

type GetAll interface {
	GetAll(skip string) (usersList []models.User, err error)
}

type GetUserById interface {
	GetUserById(id int) (user models.User, err error)
}

type CountUsers interface {
	CountUsers() (numberOfUsers int, err error)
}

type CheckIfUserExists interface {
	CheckIfUserExist(id int) error
}
