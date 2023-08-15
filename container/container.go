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

func NewContainer(config *configs.Config, DB *sqlx.DB, logging *logrus.Logger) *Container {

	return &Container{Config: config,
		DB:         DB,
		Logging:    logging,
		Repository: repositories.NewUsersRepository(DB),
	}
}
