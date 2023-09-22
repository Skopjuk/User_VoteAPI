package container

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"userapi/configs"
	"userapi/repositories"
)

type Container struct {
	Config           configs.Config
	DB               *sqlx.DB
	Logging          *logrus.Logger
	UsersRepository  *repositories.UsersRepository
	RatingRepository *repositories.RatingRepository
	VotesRepository  *repositories.VotesRepository
}

func NewContainer(config configs.Config, logging *logrus.Logger) *Container {
	db, err := NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	return &Container{Config: config,
		DB:               db,
		Logging:          logging,
		UsersRepository:  repositories.NewUsersRepository(db),
		RatingRepository: repositories.NewRatingRepository(db),
		VotesRepository:  repositories.NewVotesRepository(db),
	}
}
