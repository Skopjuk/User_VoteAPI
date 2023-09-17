package user

type DeleteRate struct {
	deleteVote DeleteVote
}

func NewDeleteRate(deleteVote DeleteVote) *DeleteRate {
	return &DeleteRate{deleteVote: deleteVote}
}

func (c *DeleteRate) Execute(id int) error {
	return c.deleteVote.DeleteVote(id)
}
