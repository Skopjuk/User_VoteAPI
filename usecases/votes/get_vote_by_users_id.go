package votes

import "userapi/usecases/user"

type GetVoteByUsersId struct {
	repository user.GetVoteByUserIds
}

func NewGetVoteByUsersId(repository user.GetVoteByUserIds) *GetVoteByUsersId {
	return &GetVoteByUsersId{repository: repository}
}

func (c *GetVoteByUsersId) Execute(id1, id2 int) (vote int, err error) {
	return c.repository.GetVoteByUserIds(id1, id2)
}
