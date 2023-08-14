package user

import (
	"github.com/sirupsen/logrus"
	"userapi/models"
)

type GetAllUsers struct {
	repository GetAll
}

func NewGetAllUsers(repository GetAll) *GetAllUsers {
	return &GetAllUsers{repository: repository}
}

func (a *GetAllUsers) Execute(skip string, paginationLimit string) ([]models.User, error) {
	logrus.Info("try to get all users")
	users, err := a.repository.GetAll(skip, paginationLimit)
	if err != nil {
		logrus.Errorf("users wern't found: %s", err)
		return nil, err
	}

	return users, nil
}
