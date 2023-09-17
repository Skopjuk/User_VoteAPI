package user

type GetUserRateById struct {
	repository GetUsersRate
}

func NewGetUserRateById(repository GetUsersRate) *GetUserRateById {
	return &GetUserRateById{repository: repository}
}

func (g *GetUserRateById) Execute(id int) (rate int, err error) {
	rate, err = g.repository.GetUsersRate(id)

	return id, err
}