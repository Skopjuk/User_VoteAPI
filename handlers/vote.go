package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"userapi/repositories"
	"userapi/usecases/user"
)

type VoteParams struct {
	UserId               int    `json:"user_id"`
	RatedUserId          int    `json:"rated_user_id"`
	UsernameWhoVotes     string `json:"username_who_votes"`
	UsernameForWhomVotes string `json:"username_for_whom_votes"`
	Rate                 int    `json:"rate"`
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

	newCheckIfUserAlreadyVoted := user.NewCheckIfUserAlreadyVotedForSomebody(v.container.Repository)
	err := newCheckIfUserAlreadyVoted.Execute(input.UserId, input.RatedUserId)
	if err != nil {
		logrus.Errorf("user with %d already voted for user with id %d", input.UserId, input.RatedUserId)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("%s", err),
		})
	}

	newGetRateByUserId := user.NewGetUserRateById(v.container.Repository)
	rate, err := newGetRateByUserId.Execute(input.UserId)
	if err != nil {
		logrus.Errorf("user with id %d wasn't find: %s", input.UserId, err)
	}

	userRepository := repositories.NewUsersRepository(v.container.DB)
	newVote := user.NewVote(userRepository)

	params := user.NewVoteAttributes{
		UserId:               input.UserId,
		RatedUserId:          input.RatedUserId,
		UsernameWhoVotes:     input.UsernameWhoVotes,
		UsernameForWhomVotes: input.UsernameForWhomVotes,
		Rate:                 rate + input.Rate,
	}
	err = newVote.Execute(params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful vote",
	})

	return err
}

func (v *VotesHandler) GetAllVotes(c echo.Context) error {
	newGetVotes := user.NewGetListOfVotes(v.container.Repository)
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
	var input user.ChangeRateAttributes

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

	newChangeVote := user.NewChangeRate(v.container.Repository)
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
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return err
	}

	newDeleteVote := user.NewDeleteRate(v.container.Repository)
	err = newDeleteVote.Execute(idInt)
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

func checkForLegitInput(rate VoteParams) error {
	if rate.Rate != 1 && rate.Rate != -1 {
		return errors.New("rate should be 1 or -1")
	} else if rate.UserId == rate.RatedUserId {
		return errors.New("user can not vote for himself")
	}

	return nil
}
