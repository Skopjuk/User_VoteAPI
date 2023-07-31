package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"userapi/models"
)

type Authenticate struct {
	repository FindUserByUsername
}

func NewAuthenticate(repository FindUserByUsername) *Authenticate {
	return &Authenticate{repository: repository}
}

type AuthenticateAttributes struct {
	Username string
	Password string
}

func (a *Authenticate) Execute(attributes AuthenticateAttributes) (*models.User, error) {
	logrus.Infof("user with username %s try to authenticate", attributes.Username)
	user, err := a.repository.FindUserByUsername(attributes.Username)
	if err != nil {
		logrus.Error("can not execute usecase: ", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(attributes.Password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
