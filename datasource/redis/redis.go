package redis

import (
	"context"

	"github.com/cpw0321/mammoth/config"
	"github.com/cpw0321/mammoth/logger"
	"github.com/go-redis/redis/v8"
)

// Redis ...
type Redis struct {
	Client *redis.Client
}

// New ...
func New(conf config.Config, log *logger.Logger) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Username: conf.Redis.Username,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Errorf("redis ping is failed, err:%v", err)
		return nil, err
	}

	return &Redis{
		Client: client,
	}, nil
}
