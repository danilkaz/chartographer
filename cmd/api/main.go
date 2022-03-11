package main

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/danilkaz/chartographer/internal/repository/storage"
	"github.com/danilkaz/chartographer/internal/service"
	"github.com/danilkaz/chartographer/internal/transport/rest"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	db := map[uuid.UUID]models.Charta{}
	storage := storage.NewStorage(&db)
	r := repository.NewRepository(storage)
	s := service.NewService(r)
	h := rest.NewHandler(s)
	err := http.ListenAndServe(":8000", h.InitRoutes())
	if err != nil {
		return
	}
}
