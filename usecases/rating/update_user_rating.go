package rating

type UpdateUsersRating struct {
	repository UpdateUserRating
}

func NewUpdateUsersRating(repository UpdateUserRating) *UpdateUsersRating {
	return &UpdateUsersRating{repository: repository}
}

func (u *UpdateUsersRating) Execute(rating, userId int) error {
	return u.repository.UpdateUserRating(rating, userId)
}
