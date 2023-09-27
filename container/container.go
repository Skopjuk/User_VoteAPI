package container

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"userapi/configs"
	"userapi/redis_db"
	"userapi/repositories"
)

type Container struct {
	Config           configs.Config
	DB               *sqlx.DB
	Logging          *logrus.Logger
	UsersRepository  *repositories.UsersRepository
	RatingRepository *repositories.RatingRepository
	VotesRepository  *repositories.VotesRepository
	RedisDb          *redis_db.RedisDb
	Context          context.Context
}

func NewContainer(config configs.Config, logging *logrus.Logger) *Container {
	db, err := NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	redisDb, err := redis_db.NewRedisDb(config.RedisHost + ":" + config.RedisPort)
	if err != nil {
		logrus.Fatalf("Failed to connect to redis: %s:", err.Error())
	}

	return &Container{Config: config,
		DB:               db,
		Logging:          logging,
		UsersRepository:  repositories.NewUsersRepository(db),
		RatingRepository: repositories.NewRatingRepository(db),
		VotesRepository:  repositories.NewVotesRepository(db),
		RedisDb:          redisDb,
		Context:          context.TODO(),
	}
}
