package main

import (
	"context"
	"fmt"
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ctx = context.Background()

func initRedisConfig() models.RedisConfig {
	return models.RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	redisConfig := initRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb)
}
