package redis_db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisDb struct {
	Client *redis.Client
}

func NewRedisDb(addr string) (*RedisDb, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &RedisDb{
		Client: client,
	}, nil
}
