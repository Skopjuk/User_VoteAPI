package user

import (
	"bytes"
	"usersAPI/models"
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
	user, err := a.repository.FindUserByUsername(attributes.Username)
	if err != nil {
		return nil, err
	}

	if bytes.Compare(user.Password, PasswordHashing(attributes.Password)) == 0 {
		return user, nil
	}

	return user, nil
}
