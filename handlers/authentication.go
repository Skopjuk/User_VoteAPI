package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"userapi/repositories"
	"userapi/usecases/user"
)

type SignUpParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

func (h *Handler) SignUp(c echo.Context) error {
	var input SignUpParams

	log := logrus.WithFields(
		logrus.Fields{
			"endpoint": "sign-up",
		})

	if err := c.Bind(&input); err != nil {
		logrus.Error("failed to bind req body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	usersRepository := repositories.NewUsersRepository(h.db)
	newProfile := user.NewCreateProfile(usersRepository)

	params := user.NewUserAttributes{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
	}
	id, err := newProfile.Execute(params)
	if err != nil {
		log.Errorf("cannot execute usecase: %s", err.Error())
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}

	return err
}

func (h *Handler) SignIn(c echo.Context) error {
	return nil
}
