package repository

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
)

type Charta interface {
	Add(charta models.Charta) (uuid.UUID, error)
	GetById(id uuid.UUID) (models.Charta, error)
	Update(id uuid.UUID, new models.Charta) error
	Delete(id uuid.UUID) error
}

type Repository struct {
	Charta
}

func NewRepository(storage map[uuid.UUID]models.Charta) *Repository {
	return &Repository{Charta: NewStorage(&storage)}
}
