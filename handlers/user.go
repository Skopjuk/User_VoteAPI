package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"userapi/models"
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

type UpdatePasswordParams struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var input UpdateUserParams

	idInt := GetUsersId(c)

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
	err := newUpdateProfile.Execute(params, idInt)
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

type QueryResult struct {
	skip  string
	limit string
}

func (h *Handler) GetAll(c echo.Context) error {
	q := c.Request().URL.Query()
	page, err := strconv.Atoi(q["page"][0])
	if err != nil {
		logrus.Errorf("error while converting page number to int: %s", err)
		return err
	}

	skip := strconv.Itoa((page - 1) * 10)
	usersRepository := repositories.NewUsersRepository(h.db)

	newGetUsers := user.NewGetAllUsers(usersRepository)
	users, err := newGetUsers.Execute(skip)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})

	return err
}

func (h *Handler) GetUserById(c echo.Context) error {
	idInt := GetUsersId(c)

	bindedUser := models.User{}
	err := c.Bind(&bindedUser)
	if err != nil {
		logrus.Error("error of binding json")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	usersRepository := repositories.NewUsersRepository(h.db)
	newGetUserById := user.NewGetUserByID(usersRepository)
	user, err := newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return err
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (h *Handler) GerNumberOfUsers(c echo.Context) error {
	usersRepository := repositories.NewUsersRepository(h.db)
	newGetUserById := user.NewCountAllUsers(usersRepository)
	numOfUsers, err := newGetUserById.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return err
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"number_of_users": numOfUsers,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (h *Handler) ChangePassword(c echo.Context) error {
	var input UpdatePasswordParams
	idInt := GetUsersId(c)

	if err := c.Bind(&input); err != nil {
		h.logging.Errorf("failedd to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	params := user.ChangePasswordAttributes{
		Password: input.Password,
	}

	usersRepository := repositories.NewUsersRepository(h.db)
	newChangePassword := user.NewChangePassword(usersRepository)
	err := newChangePassword.Execute(idInt, params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	err = c.JSON(http.StatusOK, map[string]interface{}{
		"status_password_changing": "changed",
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func GetUsersId(c echo.Context) int {
	id := c.Param("id")
	logrus.Infof("try to get user with id %s", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error of converting id to int. id: %s", id)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return idInt
}
