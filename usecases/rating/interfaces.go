package rating

import "userapi/models"

type AddUserRating interface {
	AddUserRating(rating models.Rating) error
}

type DeleteUserRating interface {
	DeleteUserRating(id int) error
}

type UpdateUserRating interface {
	UpdateUserRating(rating, id int) error
}

type GetRatingByUserId interface {
	GetRatingByUserId(id int) (rating int, err error)
}
