package votes

import "userapi/usecases/user"

type DeleteUsersVote struct {
	deleteVote user.DeleteVote
}

type DeleteRateAttributes struct {
	RatedUserId int `json:"rated_user_id"`
}

func NewDeleteUsersVote(deleteVote user.DeleteVote) *DeleteUsersVote {
	return &DeleteUsersVote{deleteVote: deleteVote}
}

func (c *DeleteUsersVote) Execute(userId, ratedUserid int) error {
	return c.deleteVote.DeleteVote(userId, ratedUserid)
}
