package rating

type DeleteRating struct {
	repository DeleteUserRating
}

func NewDeleteRating(repository DeleteUserRating) *DeleteRating {
	return &DeleteRating{repository: repository}
}

func (d *DeleteRating) Execute(id int) error {
	return d.repository.DeleteUserRating(id)
}
