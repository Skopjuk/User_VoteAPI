package user

type DeleteProfile struct {
	deleteRepository DeleteUser
	getUserById      GetUserById
}

func NewDeleteProfile(deleteRepository DeleteUser) *DeleteProfile {
	return &DeleteProfile{deleteRepository: deleteRepository}
}

func (c *DeleteProfile) Execute(id int) error {
	//authenticated := c.authenticate.AuthenticateUser(attributes.Username, attributes.Password)
	//if !authenticated {
	//	return fmt.Errorf("user is not autheniticated")
	//}

	return c.deleteRepository.DeleteUser(id)
}
