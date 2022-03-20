package main

import (
	"context"
	"github.com/danilkaz/chartographer/internal"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/danilkaz/chartographer/internal/service"
	"github.com/danilkaz/chartographer/internal/transport/rest"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func init() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config does not exist")
	}
}

func main() {
	port := viper.GetString("port")

	wd, err := os.Getwd()
	if err != nil {
		return
	}
	path := filepath.Join(wd, filepath.Dir(os.Args[1]))

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Unable to create directory %s", path)
	}

	r := repository.NewRepository(path)
	s := service.NewService(r)
	h := rest.NewHandler(s)

	server := internal.NewServer()

	go func() {
		if err = server.Run(port, h.InitRoutes()); err != nil {
			log.Fatal("Server run error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = server.Stop(ctx); err != nil {
		log.Fatal("Server shutdown error")
	}
}
