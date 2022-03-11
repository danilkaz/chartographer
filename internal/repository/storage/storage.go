package storage

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
)

type Storage struct {
	db *map[uuid.UUID]models.Charta
}

func NewStorage(storage *map[uuid.UUID]models.Charta) *Storage {
	return &Storage{db: storage}
}

func (s *Storage) Create(width, height int) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (s *Storage) GetById(id uuid.UUID) (models.Charta, error) {
	var c models.Charta
	return c, nil
}

func (s *Storage) Update(id uuid.UUID, new models.Charta) error {
	return nil
}

func (s *Storage) Delete(id uuid.UUID) error {
	return nil
}
