package user

import (
	"github.com/sirupsen/logrus"
	"userapi/models"
)

type GetUserByID struct {
	repository GetUserById
}

func NewGetUserByID(repository GetUserById) *GetUserByID {
	return &GetUserByID{repository: repository}
}

func (g *GetUserByID) Execute(id int) (user models.User, err error) {
	user, err = g.repository.GetUserById(id)
	if err != nil {
		logrus.Errorf("user with id %s not found", id)
	}

	return user, err
}
