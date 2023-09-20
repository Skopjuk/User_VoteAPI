package rating

type CreatingOrUpdatingRating struct {
	createUserRating *CreateUserRating
	updateRating     *UpdateUsersRating
	getUsersRating   *GetUserRating
}

func NewCreatingOrUpdatingRating(createUserRating *CreateUserRating, updateRating *UpdateUsersRating, getUsersRating *GetUserRating) *CreatingOrUpdatingRating {
	return &CreatingOrUpdatingRating{createUserRating: createUserRating, updateRating: updateRating, getUsersRating: getUsersRating}
}

func (c *CreatingOrUpdatingRating) Execute(userId, ratedUserId, vote int) (newRating int, err error) {
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
