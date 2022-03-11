package service

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/google/uuid"
)

type ChartaService struct {
	repository.Charta
}

func NewChartaService(repo *repository.Repository) *ChartaService {
	return &ChartaService{repo}
}

func (s *ChartaService) Create(width, height int) (uuid.UUID, error) {
	var c uuid.UUID
	return c, nil
}

func (s *ChartaService) SaveRestoredFragment(x, y, width, height int) error {
	return nil
}

func (s *ChartaService) GetPart(x, y, width, height int) (models.Charta, error) {
	var c models.Charta
	return c, nil
}

func (s *ChartaService) Delete() error {
	return nil
}
