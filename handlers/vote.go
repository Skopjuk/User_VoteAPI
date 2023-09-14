package handlers

import (
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

	userRepository := repositories.NewUsersRepository(v.container.DB)
	newVote := user.NewVote(userRepository)

	params := user.NewVoteAttributes{
		UserId:               input.UserId,
		RatedUserId:          input.RatedUserId,
		UsernameWhoVotes:     input.UsernameWhoVotes,
		UsernameForWhomVotes: input.UsernameForWhomVotes,
		Rate:                 input.Rate,
	}
	err := newVote.Execute(params)
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
