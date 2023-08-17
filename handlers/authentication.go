package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
	"userapi/repositories"
	"userapi/usecases/user"
)

const (
	signingKey = "djf2kj9(9e9)#j"
	timeTTL    = 24
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

type SignUpParams struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
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

	usersRepository := repositories.NewUsersRepository(u.container.DB)
	newProfile := user.NewCreateProfile(usersRepository)

	params := user.NewUserAttributes{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      input.Role,
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

func (u *UsersHandler) SignIn(c echo.Context) error {
	var input SignInParams
	logrus.Infof("user %s tries to authenticate", input)

	if err := c.Bind(&input); err != nil {
		u.container.Logging.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	params := user.AuthenticateAttributes{
		Username: input.Username,
		Password: input.Password,
	}

	token, err := u.GenerateToken(params.Username, params.Password)

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (u *UsersHandler) GenerateToken(username, password string) (string, error) {
	params := user.AuthenticateAttributes{
		Username: username,
		Password: password,
	}
	usersRepository := repositories.NewUsersRepository(u.container.DB)
	newAuthentication := user.NewAuthenticate(usersRepository)
	foundUser, err := newAuthentication.Execute(params)
	if err != nil {
		logrus.Errorf("cannot execute usecase: %s", err.Error())

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		UserId: foundUser.Id,
		Role:   foundUser.Role,
	})

	signedString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
