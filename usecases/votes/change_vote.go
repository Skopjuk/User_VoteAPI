package votes

import (
	"time"
	"userapi/models"
	"userapi/usecases/user"
)

type ChangeUsersVote struct {
	changeVote user.ChangeVote
}

func NewChangeVote(changeVote user.ChangeVote) *ChangeUsersVote {
	return &ChangeUsersVote{changeVote: changeVote}
}

type ChangeRateAttributes struct {
	UserId      int `json:"user_id"`
	RatedUserId int `json:"rated_user_id"`
	Vote        int `json:"vote"`
}

func (c *ChangeUsersVote) Execute(attributes ChangeRateAttributes, id int) error {
	return c.changeVote.ChangeVote(models.Votes{
		UserId:      attributes.UserId,
		RatedUserId: attributes.RatedUserId,
		Vote:        attributes.Vote,
		UpdatedAt:   time.Now(),
	}, id)
}
