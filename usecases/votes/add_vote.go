package votes

import (
	"fmt"
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
	err := v.repository.AddVoteRecord(models.Votes{
		UserId:      attributes.UserId,
		RatedUserId: attributes.RatedUserId,
		Vote:        attributes.Vote,
	})
	if err != nil {
		fmt.Errorf("can not add record to record list")
		return err
	}

	return nil
}
