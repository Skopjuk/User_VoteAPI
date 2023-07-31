package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"userapi/repositories"
	"userapi/usecases/user"
)

type UpdateUserParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type GetAllUsersParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var input UpdateUserParams

	//нужно принимать айдишник сначала і проверять его!!!

	if err := c.Bind(&input); err != nil {
		h.logging.Errorf("failedd to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	params := user.UpdateUserAttributes{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	usersRepository := repositories.NewUsersRepository(h.db)
	newUpdateProfile := user.NewChangeProfile(usersRepository)
	err := newUpdateProfile.Execute(params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return err
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"user": input.Username,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (h *Handler) GetAll(c echo.Context) error {
	usersRepository := repositories.NewUsersRepository(h.db)

	newGetUsers := user.NewGetAllUsers(usersRepository)
	users, err := newGetUsers.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})

	return err
}
