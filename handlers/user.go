package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"userapi/models"
	"userapi/usecases/user"
)

const paginationLimit = "10"

type GetAllUsersParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
}

type QueryResult struct {
	skip  string
	limit string
}

type UpdatePasswordParams struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (a *AuthHandler) UpdateUser(c echo.Context) error {
	var input user.UpdateUserAttributes

	idInt, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
		return err
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	newGetUserById := user.NewGetUserByID(a.container.Repository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("user with id %d wasn't find", idInt)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("user with id %d wasn't find", idInt),
		})
		return err
	}

	newUpdateProfile := user.NewChangeProfile(a.container.Repository)
	err = newUpdateProfile.Execute(input, idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
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

func (u *UsersHandler) GetAll(c echo.Context) error {
	pageNum := c.QueryParam("page")
	if pageNum == "" {
		pageNum = "1"
	}

	page, err := strconv.Atoi(pageNum)
	if err != nil {
		logrus.Errorf("error while converting page number to int: %s", err)
		return err
	}

	skip := strconv.Itoa((page - 1) * 10)

	newGetUsers := user.NewGetAllUsers(u.container.Repository)
	users, err := newGetUsers.Execute(skip, paginationLimit)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "error while parsing url",
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})

	return err
}

func (u *UsersHandler) GetUserById(c echo.Context) error {
	idInt, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
		return err
	}

	input := models.User{}
	err = c.Bind(&input)
	if err != nil {
		logrus.Error("error of binding json")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	newGetUserById := user.NewGetUserByID(u.container.Repository)
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

func (u *UsersHandler) GerNumberOfUsers(c echo.Context) error {
	newGetUserById := user.NewCountAllUsers(u.container.Repository)
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

func (a *AuthHandler) ChangePassword(c echo.Context) error {
	var input UpdatePasswordParams

	idInt, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
		return err
	}

	newGetUserById := user.NewGetUserByID(a.container.Repository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("user with id %d wasn't find", idInt)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("user with id %d wasn't find", idInt),
		})
		return err
	}

	if err := c.Bind(&input); err != nil {
		a.container.Logging.Errorf("failedd to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	params := user.ChangePasswordAttributes{
		Password: input.Password,
	}

	newChangePassword := user.NewChangePassword(a.container.Repository)
	err = newChangePassword.Execute(idInt, params)
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

func getUserId(c echo.Context) (int, error) {
	id := c.Param("id")
	logrus.Infof("try to get user with id %s", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error of converting id to int. id: %s", id)
		c.JSON(http.StatusInternalServerError, err.Error())
		return 0, err
	}

	return idInt, nil
}
