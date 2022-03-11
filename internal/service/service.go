package service

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
)

type Service struct {
	models.ChartaInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo}
}
