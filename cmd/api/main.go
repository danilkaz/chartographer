package main

import (
	"fmt"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	rdb := repository.NewRedisDatabase(repository.RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	})
	fmt.Println(rdb)
}
