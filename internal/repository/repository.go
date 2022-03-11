package repository

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
)

type Charta interface {
	Create(width, height int) (uuid.UUID, error)
	GetById(id uuid.UUID) (models.Charta, error)
	Update(id uuid.UUID, new models.Charta) error
	Delete(id uuid.UUID) error
}

type Repository struct {
	Charta
}

func NewRepository(storage *Storage) *Repository {
	return &Repository{Charta: storage}
}
