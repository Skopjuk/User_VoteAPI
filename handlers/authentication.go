package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"userapi/configs"
	"userapi/repositories"
	"userapi/usecases/user"
)

const (
	timeTTL = 24
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

func userIdentity(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "authorization header is empty",
		})
		logrus.Errorf("authorization header is empty, user: %s", c.Param("username"))
		return errors.New("authorization header is empty")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "authorization header is invalid",
		})
		logrus.Errorf("authorization header is invalid, user: %s", c.Param("username"))
		return errors.New("authorization header is invalid")
	}

	userId, userRole, err := parseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		logrus.Errorf("token can not be parsed")
		return errors.New(err.Error())
	}

	c.Set("userId", userId)
	c.Set("userRole", userRole)

	return nil
}

func parseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		config, err := configs.NewConfig()
		if err != nil {
			logrus.Error("config is not available")
			return "", err
		}

		return []byte(config.SigningKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, claims.Role, nil

}
