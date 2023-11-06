package votes

import (
	"errors"
	"github.com/sirupsen/logrus"
	"userapi/models"
	"userapi/usecases/user"
)

type Vote struct {
	repository user.AddVoteRecord
}

type NewVoteAttributes struct {
	UserId      int
	RatedUserId int
	Vote        int
}

func NewVote(repository user.AddVoteRecord) *Vote {
	return &Vote{repository: repository}
}

func (v *Vote) Execute(attributes NewVoteAttributes) error {
	err := paramsValidation(attributes.Vote)
	if err != nil {
		return err
	}

	err = v.repository.AddVoteRecord(models.Votes{
		UserId:      attributes.UserId,
		RatedUserId: attributes.RatedUserId,
		Vote:        attributes.Vote,
	})
	if err != nil {
		logrus.Errorf("can not add record to record list")
	}

	return err
}

func paramsValidation(vote int) error {
	if vote != 1 && vote != -1 {
		return errors.New("vote is not equal to 1 or -1")
	}

	return nil
}