package container

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"userapi/configs"
	"userapi/repositories"
)

type Container struct {
	Config     *configs.Config
	DB         *sqlx.DB
	Logging    *logrus.Logger
	Repository *repositories.UsersRepository
}
