package rating

type CreateOrUpdateRating struct {
	createUserRating *CreateUserRating
	updateRating     *UpdateUsersRating
	getUsersRating   *GetUserRating
}

func NewCreateOrUpdateRating(createUserRating *CreateUserRating, updateRating *UpdateUsersRating, getUsersRating *GetUserRating) *CreateOrUpdateRating {
	return &CreateOrUpdateRating{createUserRating: createUserRating, updateRating: updateRating, getUsersRating: getUsersRating}
}

func (c *CreateOrUpdateRating) Execute(ratedUserId, vote int) (newRating int, err error) {
	rating, err := c.getUsersRating.Execute(ratedUserId)
	if err != nil {
		err := c.createUserRating.Execute(UsersRatingAttributes{
			UserId: ratedUserId,
			Rating: vote,
		})
		if err != nil {
			return 0, err
		}
	} else {
		newRating := rating + vote
		err := c.updateRating.Execute(newRating, ratedUserId)
		if err != nil {
			return 0, err
		}

	}
	return newRating, nil
}
