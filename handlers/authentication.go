package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	var input *SignUpParams

	log := logrus.WithFields(
		logrus.Fields{
			"endpoint": "sign-up",
		})

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Errorf("cannot connect to db: %s", err.Error())
		return err
	}

	if err := c.Bind(&input); err != nil {
		logrus.Error("failed to bind req body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	usersRepository := repositories.NewUsersRepository(db)
	newProfile := user.NewCreateProfile(usersRepository)
	err = c.Bind(&input)
	if err != nil {
		log.Errorf("cannot parse query: %s", err.Error())
		return err
	}

	params := user.NewUserAttributes{
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
	}
	err = newProfile.Execute(params)
	if err != nil {
		log.Error("cannot execute usecase: %s", err.Error())
		return err
	}
	return nil
}

func (h *Handler) SignIn(c echo.Context) error {
	return nil
}
