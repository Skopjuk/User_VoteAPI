package user

import (
	"fmt"
)

type ChangePassword struct {
	repository          AuthenticateUser
	changeUsersPassword ChangeUsersPassword
}

func NewChangePassword() *ChangePassword {
	return &ChangePassword{}
}

type ChangePasswordAttributes struct {
	Username string
	Password string
}

func (a *ChangePassword) Execute(attributes ChangePasswordAttributes) (bool, error) {
	authenticated := a.repository.AuthenticateUser(attributes.Username, attributes.Password)
	if !authenticated {
		return false, fmt.Errorf("user is not authenticated")
	}

	return a.changeUsersPassword.ChangeUsersPassword(attributes.Username, attributes.Password), nil
}
