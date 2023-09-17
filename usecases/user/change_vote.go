package user

import (
	"time"
	"userapi/models"
)

type ChangeRate struct {
	changeVote ChangeVote
}

func NewChangeRate(changeVote ChangeVote) *ChangeRate {
	return &ChangeRate{changeVote: changeVote}
}

type ChangeRateAttributes struct {
	UserId               int    `json:"user_id"`
	RatedUserId          int    `json:"rated_user_id"`
	UsernameWhoVotes     string `json:"username_who_votes"`
	UsernameForWhomVotes string `json:"username_for_whom_votes"`
	Rate                 int    `json:"rate"`
}

func (c *ChangeRate) Execute(attributes ChangeRateAttributes, id int) error {
	return c.changeVote.ChangeVote(models.Rate{
		UserId:               attributes.UserId,
		RatedUserId:          attributes.RatedUserId,
		UsernameWhoVotes:     attributes.UsernameWhoVotes,
		UsernameForWhomVotes: attributes.UsernameForWhomVotes,
		Rate:                 attributes.Rate,
		UpdatedAt:            time.Now(),
	}, id)
}
