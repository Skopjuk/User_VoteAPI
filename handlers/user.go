package handlers

import (
	"encoding/json"
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

func (a *AccountHandler) UpdateUser(c echo.Context) error {
	var input user.UpdateUserAttributes

	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	newGetUserById := user.NewGetUserByID(a.container.UsersRepository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem wile inserting user: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem wile inserting user: %s", err),
		})
	}

	newUpdateProfile := user.NewChangeProfile(a.container.UsersRepository)
	err = newUpdateProfile.Execute(input, idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
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
	var input []models.User
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

	val, err := u.container.RedisDb.Get(c.Request().Context(), "users").Result()
	if val != "" {
		logrus.Info("data about users list exists in redis")
		if err := json.Unmarshal([]byte(val), &input); err != nil {
			logrus.Errorf("failed to bind req body: %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user": input,
		})
	}

	logrus.Info("in redis no data about all users. Request to Postrgres")
	newGetUsers := user.NewGetAllUsers(u.container.UsersRepository)
	users, err := newGetUsers.Execute(skip, paginationLimit)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "error while parsing url",
		})
	}

	var usersList []models.User
	for _, i := range users {
		foundUser := i
		usersList = append(usersList, foundUser)
	}

	data, err := json.Marshal(usersList)
	if err != nil {
		logrus.Errorf("error while marshaling users list")
		return err
	}

	u.container.RedisDb.Set(c.Request().Context(), "users", data, u.container.Config.ExpTime)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": users,
	})
}

func (u *UsersHandler) GetUserById(c echo.Context) error {
	var input models.User
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	keyForRedis := "user_by_id_" + strconv.Itoa(idInt)

	userById, err := u.container.RedisDb.Get(c.Request().Context(), keyForRedis).Result()
	if err != nil {
		logrus.Errorf("error while getting data from redis")
		return c.JSON(http.StatusInternalServerError, err)
	}

	if userById != "" {
		logrus.Info("data about this user exists in redis")
		if err := json.Unmarshal([]byte(userById), &input); err != nil {
			logrus.Errorf("failed to bind req body: %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user": input,
		})
	}

	logrus.Info("in redis no data about this user. Request to Postrgres")

	newGetUserById := user.NewGetUserByID(u.container.UsersRepository)
	foundUser, err := newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	u.container.RedisDb.Set(c.Request().Context(), keyForRedis, foundUser, u.container.Config.ExpTime)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": foundUser,
	})
}

func (u *UsersHandler) GetNumberOfUsers(c echo.Context) error {
	usersNumRedis, err := u.container.RedisDb.Get(c.Request().Context(), "amount_of_users").Result()

	if usersNumRedis != "" {
		logrus.Info("data about amount of users exists in redis")

		usersNumRedisStr, err := strconv.Atoi(usersNumRedis)
		if err != nil {
			logrus.Errorf("error while converting amount of users to string: %s", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"number_of_users": usersNumRedisStr,
		})
	}

	logrus.Info("in redis no data about amount of users. Request to Postrgres")
	newGetUserById := user.NewCountAllUsers(u.container.UsersRepository)
	numOfUsers, err := newGetUserById.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	u.container.RedisDb.Set(c.Request().Context(), "amount_of_users", numOfUsers, u.container.Config.ExpTime)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"number_of_users": numOfUsers,
	})
}

func (a *AccountHandler) ChangePassword(c echo.Context) error {
	var input UpdatePasswordParams

	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	newGetUserById := user.NewGetUserByID(a.container.UsersRepository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem wile inserting user: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem wile inserting user: %s", err),
		})
	}

	if err := c.Bind(&input); err != nil {
		a.container.Logging.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	params := user.ChangePasswordAttributes{
		Password: input.Password,
	}

	newChangePassword := user.NewChangePassword(a.container.UsersRepository)
	err = newChangePassword.Execute(idInt, params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status_password_changing": "changed",
	})
}

func (a *AccountHandler) DeleteUser(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return err
	}

	newGetUserById := user.NewGetUserByID(a.container.UsersRepository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem wile inserting user: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem wile inserting user: %s", err),
		})
	}

	newDeleteUser := user.NewDeleteProfile(a.container.UsersRepository)
	err = newDeleteUser.Execute(idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status_deleting_user": "deleted",
	})
}

func getIdFromEndpoint(c echo.Context) (int, error) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error of converting id to int. id: %s", id)
		return 0, c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("id %d can not be parsed", idInt),
		})
	}

	return idInt, nil
}
