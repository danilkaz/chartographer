package repository

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

func (s *Storage) SaveRestoredFragment(x, y, width, height int) error {
	return nil
}

func (s *Storage) GetPart(x, y, width, height int) (models.Charta, error) {
	var c models.Charta
	return c, nil
}

func (s *Storage) Delete() error {
	return nil
}
