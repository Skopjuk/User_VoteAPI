package user

import (
	"fmt"
	"userapi/models"
)

//позбувитись від бібліотеки errors на користь fmt.Errorf

type ChangeProfile struct {
	updateUser  UpdateUser
	getUserById GetUserById
}

func NewChangeProfile(updateUser UpdateUser) *ChangeProfile {
	return &ChangeProfile{updateUser: updateUser}
}

type UpdateUserAttributes struct {
	Username  string
	FirstName string
	LastName  string
}

func (c *ChangeProfile) Execute(attributes UpdateUserAttributes, id int) error {
	//винести перевірки в окрему функцію
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

	return c.updateUser.UpdateUser(models.User{
		Username:  attributes.Username,
		FirstName: attributes.FirstName,
		LastName:  attributes.LastName,
	}, id)
}
