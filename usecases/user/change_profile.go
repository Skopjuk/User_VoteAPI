package user

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"userapi/models"
)

//позбувитись від бібліотеки errors на користь fmt.Errorf

type ChangeProfile struct {
	updateUser UpdateUser
}

func NewChangeProfile(updateUser UpdateUser) *ChangeProfile {
	return &ChangeProfile{updateUser: updateUser}
}

type UpdateUserAttributes struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
}

func (c *ChangeProfile) Execute(attributes UpdateUserAttributes, id int) error {
	err := validateUser(attributes)
	if err != nil {
		logrus.Errorf("error while updating user: %s", err)
		return err
	}

	return c.updateUser.UpdateUser(models.User{
		Username:  attributes.Username,
		FirstName: attributes.FirstName,
		LastName:  attributes.LastName,
		Role:      attributes.Role,
	}, id)
}

func validateUser(attributes UpdateUserAttributes) error {
	if len(attributes.FirstName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.FirstName) > 50 {
		return fmt.Errorf("first name is too long")
	} else if len(attributes.LastName) < 2 {
		return fmt.Errorf("first name is too short")
	} else if len(attributes.LastName) > 50 {
		return fmt.Errorf("last name is too long")
	} else if len(attributes.Username) < 3 {
		return fmt.Errorf("username is too short")
	} else if len(attributes.Username) > 30 {
		return fmt.Errorf("username is too long")
	}

	return nil
}
