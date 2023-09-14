package user

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"userapi/models"
)

type Vote struct {
	repository AddVoteRecord
}

type NewVoteAttributes struct {
	UserId               int
	RatedUserId          int
	UsernameWhoVotes     string
	UsernameForWhomVotes string
	Rate                 int
}

func NewVote(repository AddVoteRecord) *Vote {
	return &Vote{repository: repository}
}

func (v *Vote) Execute(attributes NewVoteAttributes) error {
	err := voteParametersValidation(attributes)
	if err != nil {
		logrus.Errorf("error while adding vote: %s", err)
		return err
	}

	err = v.repository.AddVoteRecord(models.Rate{
		UserId:               attributes.UserId,
		RatedUserId:          attributes.RatedUserId,
		UsernameWhoVotes:     attributes.UsernameWhoVotes,
		UsernameForWhomVotes: attributes.UsernameForWhomVotes,
		Rate:                 attributes.Rate,
	})
	if err != nil {
		fmt.Errorf("can not add record to record list")
		return err
	}

	return nil
}

func voteParametersValidation(attributes NewVoteAttributes) error {
	if attributes.Rate != 1 && attributes.Rate != -1 {
		return errors.New("rate should be -1 or 1")
	}
	return nil
}
