package user

import (
	"github.com/sirupsen/logrus"
	"userapi/models"
)

type GetListOfVotes struct {
	repository GetAllVotes
}

func NewGetListOfVotes(repository GetAllVotes) *GetListOfVotes {
	return &GetListOfVotes{repository: repository}
}

func (g *GetListOfVotes) Execute() ([]models.Rate, error) {
	logrus.Info("try to get all votes")
	votes, err := g.repository.GetAllVotes()
	if err != nil {
		logrus.Errorf("votes wern't found: %s", err)
		return nil, err
	}

	return votes, nil
}
