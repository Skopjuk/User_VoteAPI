package user

type CheckIfUserAlreadyVotedForSomebody struct {
	repository CheckIfUserVotedForSomeUser
}

func NewCheckIfUserAlreadyVotedForSomebody(repository CheckIfUserVotedForSomeUser) *CheckIfUserAlreadyVotedForSomebody {
	return &CheckIfUserAlreadyVotedForSomebody{repository: repository}
}

func (c *CheckIfUserAlreadyVotedForSomebody) Execute(id1, id2 int) (err error) {
	err = c.repository.CheckIfUserVotedForSomeUser(id1, id2)

	return err
}
