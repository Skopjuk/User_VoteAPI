package rating

type GetUserRating struct {
	repository GetRatingByUserId
}

func NewGetUserRating(repository GetRatingByUserId) *GetUserRating {
	return &GetUserRating{repository: repository}
}

func (g *GetUserRating) Execute(userId int) (rating int, err error) {
	return g.repository.GetRatingByUserId(userId)
}
