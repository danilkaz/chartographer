package main

import (
	"fmt"
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func initDatabaseConfig() models.DatabaseConfig {
	return models.DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}
}

func openDatabaseConnection(config models.DatabaseConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)
	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	databaseConfig := initDatabaseConfig()
	databaseConnection, err := openDatabaseConnection(databaseConfig)
	if err != nil {
		log.Fatal("Database connection error")
	}
	fmt.Println(databaseConnection)
}
