package votes

import (
	"github.com/sirupsen/logrus"
	"userapi/models"
	"userapi/usecases/user"
)

type GetListOfVotes struct {
	repository user.GetAllVotes
}

func NewGetListOfVotes(repository user.GetAllVotes) *GetListOfVotes {
	return &GetListOfVotes{repository: repository}
}

func (g *GetListOfVotes) Execute() ([]models.Votes, error) {
	logrus.Info("try to get all votes")
	votes, err := g.repository.GetAllVotes()
	if err != nil {
		logrus.Errorf("votes wern't found: %s", err)
		return nil, err
	}

	return votes, nil
}
