package user

import (
	"github.com/sirupsen/logrus"
)

type CountAllUsers struct {
	repository CountUsers
}

func NewCountAllUsers(repository CountUsers) *CountAllUsers {
	return &CountAllUsers{repository: repository}
}

func (a *CountAllUsers) Execute() (int, error) {
	logrus.Info("try to get all users")
	numberOfUsers, err := a.repository.CountUsers()
	if err != nil {
		logrus.Errorf("users can not be countet: %s", err)
	}

	return numberOfUsers, err
}
