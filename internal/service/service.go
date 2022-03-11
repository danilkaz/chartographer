package service

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/google/uuid"
)

type Charta interface {
	Create(width, height int) (uuid.UUID, error)
	SaveRestoredFragment(x, y, width, height int) error
	GetPart(x, y, width, height int) (models.Charta, error)
	Delete() error
}

type Service struct {
	Charta
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Charta: NewChartaService(repo)}
}
