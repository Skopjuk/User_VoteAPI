package user

type DeleteProfile struct {
	deleteRepository DeleteUser
}

func NewDeleteProfile(deleteRepository DeleteUser) *DeleteProfile {
	return &DeleteProfile{deleteRepository: deleteRepository}
}

func (c *DeleteProfile) Execute(id int) error {
	return c.deleteRepository.DeleteUser(id)
}
