package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userapi/configs"
	"userapi/repositories"
	"userapi/usecases/user"
)

type SignInParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginHandler) SignIn(c echo.Context) error {
	var input SignInParams
	logrus.Infof("user %s tries to authenticate", input)

	if err := c.Bind(&input); err != nil {
		l.container.Logging.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	params := user.AuthenticateAttributes{
		Username: input.Username,
		Password: input.Password,
	}

	tokenRedis, err := l.container.RedisDb.Client.Get(l.container.Context, "users_token").Result()
	if err == redis.Nil {
		token, err := l.GenerateToken(params.Username, params.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "user do not exist",
			})
		}

		err = c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
		if err != nil {
			logrus.Errorf("troubles with sending http status: %s", err)
			err = c.JSON(http.StatusInternalServerError, err)
		}
		l.container.RedisDb.Client.Set(l.container.Context, "users_token", token, 0)
		_, err = l.container.RedisDb.Client.Expire(l.container.Context, "users_token", 1*time.Minute).Result()
		if err != nil {
			logrus.Errorf("error while set expirational perion")
		}
	} else if err != nil {
		logrus.Errorf("error while attempt to recive data about users rating")
	} else {
		logrus.Info("data about users token exists in redis")
		err = c.JSON(http.StatusOK, map[string]interface{}{
			"token": tokenRedis,
		})
		if err != nil {
			logrus.Errorf("troubles with sending http status: %s", err)
		}
	}

	return err
}

func (l *LoginHandler) GenerateToken(username, password string) (string, error) {
	params := user.AuthenticateAttributes{
		Username: username,
		Password: password,
	}
	usersRepository := repositories.NewUsersRepository(l.container.DB)
	newAuthentication := user.NewAuthenticate(usersRepository)
	foundUser, err := newAuthentication.Execute(params)
	if err != nil {
		logrus.Errorf("cannot execute usecase: %s", err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * timeTTL)),
		},
		UserId: foundUser.Id,
		Role:   foundUser.Role,
	})

	config, err := configs.NewConfig()
	if err != nil {
		logrus.Error("config is not available")
		return "", err
	}
	signedString, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
