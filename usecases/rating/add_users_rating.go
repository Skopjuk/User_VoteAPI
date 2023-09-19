package rating

import (
	"userapi/models"
)

type CreateUserRating struct {
	repository AddUserRating
}

func NewCreateUserRating(repository AddUserRating) *CreateUserRating {
	return &CreateUserRating{repository: repository}
}

type UsersRatingAttributes struct {
	UserId int
	Rating int
}

func (c *CreateUserRating) Execute(attributes UsersRatingAttributes) error {
	err := c.repository.AddUserRating(models.Rating{
		UserId: attributes.UserId,
		Rating: attributes.Rating,
	})

	return err
}
