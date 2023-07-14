package user

import "fmt"

type DeleteProfile struct {
	deleteRepository DeleteUserRepository
	authenticate     AuthenticateUser
}

func NewDeleteProfile() *DeleteProfile {
	return &DeleteProfile{}
}

type DeleteUserRepository interface {
	FindUserByUsername
	AuthenticateUser
	DeleteUser
}

type DeleteUserAttributes struct {
	Username string
	Password string
}

func (c *DeleteProfile) Execute(attributes DeleteUserAttributes) error {
	authenticated := c.authenticate.AuthenticateUser(attributes.Username, attributes.Password)
	if !authenticated {
		return fmt.Errorf("user is not autheniticated")
	}

	return c.deleteRepository.DeleteUser(attributes.Username)
}
