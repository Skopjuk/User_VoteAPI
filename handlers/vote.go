package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userapi/repositories"
	"userapi/usecases/rating"
	"userapi/usecases/votes"
)

type VoteParams struct {
	UserId      int `json:"user_id"`
	RatedUserId int `json:"rated_user_id"`
	Vote        int `json:"vote"`
}

func (v *VotesHandler) Vote(c echo.Context) error {
	var input VoteParams

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := checkForLegitInput(input); err != nil {
		logrus.Errorf("error: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("%s", err),
		})
	}

	err := v.checkIfUserCanVote(input.UserId, input.RatedUserId, c)
	if err != nil {
		return err
	}

	err = v.vote(input.UserId, input.RatedUserId, input.Vote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = v.checkIfUserAlreadyHaveRating(input.RatedUserId)
	if err != nil {
		err = v.createUserRating(input.RatedUserId, input.Vote)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "error while listing users rating",
			})
		}
	} else {
		newGetUserRating := rating.NewGetUserRating(v.container.RatingRepository)
		userRating, err := newGetUserRating.Execute(input.RatedUserId)
		newRating := userRating + input.Vote

		err = v.updateUsersRating(input.RatedUserId, newRating)
		if err != nil {
			logrus.Errorf("Error while updating users rating: %s", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful vote",
	})

	return err
}

func (v *VotesHandler) GetAllVotes(c echo.Context) error {
	newGetVotes := votes.NewGetListOfVotes(v.container.VotesRepository)
	votes, err := newGetVotes.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, "")
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"votes": votes,
	})
	return err
}

func (v *VotesHandler) UpdateVote(c echo.Context) error {
	var input votes.ChangeRateAttributes

	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("id %d can not be parsed", idInt),
		})
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error with parsing request body",
		})
	}

	newChangeVote := votes.NewChangeVote(v.container.VotesRepository)
	err = newChangeVote.Execute(input, idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"vote": input,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (v *VotesHandler) DeleteVote(c echo.Context) error {
	var input votes.DeleteRateAttributes
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return err
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error with parsing request body",
		})
	}

	newGetVoteByUsersId := votes.NewGetVoteByUsersId(v.container.VotesRepository)
	vote, _ := newGetVoteByUsersId.Execute(idInt, input.RatedUserId)

	newGetUserRating := rating.NewGetUserRating(v.container.RatingRepository)
	userRating, err := newGetUserRating.Execute(input.RatedUserId)
	newRating := userRating - vote
	err = v.updateUsersRating(input.RatedUserId, newRating)

	newDeleteVote := votes.NewDeleteUsersVote(v.container.VotesRepository)
	err = newDeleteVote.Execute(idInt, input.RatedUserId)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"status_deleting_vote": "success",
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func checkForLegitInput(voteParams VoteParams) error {
	if voteParams.Vote != 1 && voteParams.Vote != -1 {
		return errors.New("vote should be 1 or -1")
	} else if voteParams.UserId == voteParams.RatedUserId {
		return errors.New("user can not vote for himself")
	}

	return nil
}

func (v *VotesHandler) checkIfUserAlreadyHaveRating(userId int) error {
	newCheckIfUserHaveRating := rating.NewGetUserRating(v.container.RatingRepository)
	_, err := newCheckIfUserHaveRating.Execute(userId)

	return err
}

func (v *VotesHandler) updateUsersRating(userId, userRating int) error {
	newUpdateRating := rating.NewUpdateUsersRating(v.container.RatingRepository)
	err := newUpdateRating.Execute(userRating, userId)

	return err
}

func (v *VotesHandler) createUserRating(userId, vote int) error {
	newCreateNewRating := rating.NewCreateUserRating(v.container.RatingRepository)
	err := newCreateNewRating.Execute(rating.UsersRatingAttributes{
		UserId: userId,
		Rating: vote,
	})

	return err
}

func (v *VotesHandler) checkIfUserCanVote(userId, RatedUserId int, c echo.Context) error {
	newCheckIfUserAlreadyVoted := votes.NewGetVoteByUsersId(v.container.VotesRepository)
	_, err := newCheckIfUserAlreadyVoted.Execute(userId, RatedUserId)
	if err != nil {
		logrus.Errorf("user with %d already voted for user with id %d", userId, RatedUserId)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("%s", err),
		})
	}

	newOneHourCheck := votes.NewFindLastVote(v.container.VotesRepository)
	updatedAt, err := newOneHourCheck.Execute(userId)
	timeDiff := time.Now().Hour() - updatedAt.Hour()
	if timeDiff < 1 {
		return c.JSON(http.StatusPreconditionFailed, map[string]interface{}{
			"error": "you can vote maximum once per hour",
		})
	}

	return nil
}

func (v *VotesHandler) vote(userId, ratedUserId, vote int) error {
	userRepository := repositories.NewVotesRepository(v.container.DB)
	newVote := votes.NewVote(userRepository)

	params := votes.NewVoteAttributes{
		UserId:      userId,
		RatedUserId: ratedUserId,
		Vote:        vote,
	}
	err := newVote.Execute(params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
	}

	return err
}
