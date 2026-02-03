package storage

import (
	"context"
	"os"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func InitRedis() (*RedisClient, error) {
	addr := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: pass,
		DB: 0,
	})

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{Client: rdb}, nil
}
