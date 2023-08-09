package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"userapi/repositories"
	"userapi/usecases/user"
)

type SignUpParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

type SignInParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UsersHandler) SignUp(c echo.Context) error {
	var input SignUpParams

	log := logrus.WithFields(
		logrus.Fields{
			"endpoint": "sign-up",
		})

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	usersRepository := repositories.NewUsersRepository(u.db)
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
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "this user already exist",
			})
		} else {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return err
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

//func (h *Handler) SignIn(c echo.Context) error {
//	var input SignInParams
//
//	if err := c.Bind(&input); err != nil {
//		h.logging.Errorf("failed to bind req body: %s", err)
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//
//	params := user.AuthenticateAttributes{
//		Username: input.Username,
//		Password: input.Password,
//	}
//
//	usersRepository := repositories.NewUsersRepository(h.db)
//	newAuthentication := user.NewAuthenticate(usersRepository)
//	foundUser, err := newAuthentication.Execute(params)
//	if err != nil {
//		logrus.Errorf("cannot execute usecase: %s", err.Error())
//		c.JSON(http.StatusUnauthorized, map[string]interface{}{
//			"user": foundUser,
//		})
//		return err
//	}
//
//	err = c.JSON(http.StatusOK, map[string]interface{}{
//		"user": foundUser,
//	})
//	if err != nil {
//		logrus.Errorf("troubles with sending http status: %s", err)
//	}
//
//	return err
//}
