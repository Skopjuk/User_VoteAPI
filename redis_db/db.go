package redis_db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisDb(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
