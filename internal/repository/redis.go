package repository

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host string
	Port string
}

func NewRedisDatabase(config RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: "",
		DB:       0,
	})
}
