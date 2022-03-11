package service

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/google/uuid"
)

type Charta interface {
	Create(width, height int) (uuid.UUID, error)
	SaveRestoredFragment(id uuid.UUID, x, y, width, height int, fragment models.Charta) error
	GetPart(id uuid.UUID, x, y, width, height int) (models.Charta, error)
	Delete(id uuid.UUID) error
}

type Service struct {
	Charta
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Charta: NewChartaService(repo)}
}
