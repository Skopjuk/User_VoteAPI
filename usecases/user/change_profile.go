package user

import (
	"fmt"
	"userapi/models"
)

//позбувитись від бібліотеки errors на користь fmt.Errorf

type ChangeProfile struct {
	repository UpdateUser
}

func NewChangeProfile(repository UpdateUser) *ChangeProfile {
	return &ChangeProfile{repository: repository}
}

type UpdateUserAttributes struct {
	Username  string
	FirstName string
	LastName  string
	Password  []byte
}

func (c *ChangeProfile) Execute(attributes UpdateUserAttributes) error {
	if len(attributes.FirstName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.FirstName) > 50 {
		return fmt.Errorf("first name is too long")
	} else if len(attributes.LastName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.LastName) > 50 {
		return fmt.Errorf("last name is too long")
	} else if len(attributes.Password) < 6 {
		return fmt.Errorf("password is too short")
	} else if len(attributes.Username) < 3 {
		return fmt.Errorf("username is too short")
	} else if len(attributes.Username) > 30 {
		return fmt.Errorf("username is too long")
	}

	return c.repository.UpdateUser(&models.User{
		Username:  attributes.Username,
		FirstName: attributes.FirstName,
		LastName:  attributes.LastName,
		Password:  attributes.Password,
	})
}
