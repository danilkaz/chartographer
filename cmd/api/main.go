package main

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/danilkaz/chartographer/internal/service"
	"github.com/danilkaz/chartographer/internal/transport/rest"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	db := map[uuid.UUID]models.Charta{}
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := rest.NewHandler(s)
	err := http.ListenAndServe(":8000", h.InitRoutes())
	if err != nil {
		return
	}
}
