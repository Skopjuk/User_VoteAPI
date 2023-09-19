package votes

import "userapi/usecases/user"

type GetUserRateById struct {
	repository user.GetUsersRate
}

func NewGetUserRateById(repository user.GetUsersRate) *GetUserRateById {
	return &GetUserRateById{repository: repository}
}

func (g *GetUserRateById) Execute(id int) (rate int, err error) {
	rate, err = g.repository.GetUsersRate(id)

	return id, err
}
