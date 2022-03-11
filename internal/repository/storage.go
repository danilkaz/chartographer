package repository

import (
	"fmt"
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"os"
	"sync"
)

type Storage struct {
	mu sync.Mutex
	db *map[uuid.UUID]models.Charta
}

func NewStorage(storage *map[uuid.UUID]models.Charta) *Storage {
	return &Storage{db: storage}
}

func (s *Storage) Add(charta models.Charta) (uuid.UUID, error) {
	id := uuid.New()
	charta.Id = id
	file, err := os.Create(fmt.Sprintf("%s.bmp", id))
	defer func(file *os.File) {
		if file.Close() != nil {
			return
		}
	}(file)
	if err != nil {
		return uuid.Nil, err
	}
	err = bmp.Encode(file, charta.Image)
	if err != nil {
		return [16]byte{}, err
	}
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
