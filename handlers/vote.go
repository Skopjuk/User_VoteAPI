package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userapi/container"
	"userapi/models"
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

	err := v.checkIfUserCanVote(input.UserId, input.RatedUserId, c)
	if err != nil {
		return err
	}

	err = v.vote(input.UserId, input.RatedUserId, input.Vote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	newCreateRating := rating.NewCreateUserRating(v.container.RatingRepository)
	newUpdateRating := rating.NewUpdateUsersRating(v.container.RatingRepository)
	newGetUserRating := rating.NewGetUserRating(v.container.RatingRepository)

	newCreatingOrUpdatingRating := rating.NewCreateOrUpdateRating(newCreateRating, newUpdateRating, newGetUserRating)
	_, err = newCreatingOrUpdatingRating.Execute(input.RatedUserId, input.Vote)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful vote",
	})

	return err
}

func (v *VotesHandler) GetAllVotes(c echo.Context) error {
	var input []models.Votes
	newGetVotes := votes.NewGetListOfVotes(v.container.VotesRepository)

	redisVotes, err := v.container.RedisDb.Get(c.Request().Context(), "votes").Result()
	if err != nil {
		logrus.Errorf("error while getting data from redis: %s", err)
	}

	if redisVotes != "" {
		logrus.Info("data about all votes list exists in redis")
		if err := json.Unmarshal([]byte(redisVotes), &input); err != nil {
			logrus.Errorf("failed to bind req body: %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"votes": input,
		})
	}

	logrus.Info("in redis no data about all votes. Request to Postrgres")
	votes, err := newGetVotes.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, "")
	}

	var votesList []models.Votes
	for _, i := range votes {
		foundVote := i
		votesList = append(votesList, foundVote)
	}

	data, err := json.Marshal(votesList)
	if err != nil {
		logrus.Errorf("error while marshaling votes list:%s", err)
	}
	v.container.RedisDb.Set(c.Request().Context(), "votes", data, v.container.Config.ExpTime)

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"votes": votes,
	})

	return err
}

func (v *VotesHandler) UpdateVote(c echo.Context) error {
	var input votes.ChangeRateAttributes

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error with parsing request body",
		})
	}

	newChangeVote := votes.NewChangeVote(v.container.VotesRepository)
	err := newChangeVote.Execute(input)
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

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error with parsing request body",
		})
	}

	newGetVoteByUsersId := votes.NewGetVoteByUsersId(v.container.VotesRepository)
	vote, err := newGetVoteByUsersId.Execute(input.UserId, input.RatedUserId)
	if err.Error() != fmt.Sprintf("user with id %d already voted for user with id %d", input.UserId, input.RatedUserId) && err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	newGetUserRating := rating.NewGetUserRating(v.container.RatingRepository)
	userRating, err := newGetUserRating.Execute(input.RatedUserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	newRating := userRating - vote
	err = UpdateUsersRating(input.RatedUserId, newRating, *v.container)

	newDeleteVote := votes.NewDeleteUsersVote(v.container.VotesRepository)
	err = newDeleteVote.Execute(input.UserId, input.RatedUserId)
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

func UpdateUsersRating(userId, userRating int, container container.Container) error {
	newUpdateRating := rating.NewUpdateUsersRating(container.RatingRepository)
	err := newUpdateRating.Execute(userRating, userId)

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
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

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
