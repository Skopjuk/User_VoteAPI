package votes

import "userapi/usecases/user"

type DeleteUsersVote struct {
	deleteVote user.DeleteVote
}

func NewDeleteUsersVote(deleteVote user.DeleteVote) *DeleteUsersVote {
	return &DeleteUsersVote{deleteVote: deleteVote}
}

func (c *DeleteUsersVote) Execute(id int) error {
	return c.deleteVote.DeleteVote(id)
}
