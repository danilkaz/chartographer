package main

import (
	"github.com/danilkaz/chartographer/internal/transport/rest"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: rest.InitRoutes(),
	}
	log.Fatal(server.ListenAndServe())
}
