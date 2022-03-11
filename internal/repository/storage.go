package repository

import (
	"errors"
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
	(*s.db)[id] = charta
	if err := s.WriteToFile(charta); err != nil {
		return [16]byte{}, err
	}
	return id, nil
}

func (s *Storage) GetById(id uuid.UUID) (models.Charta, error) { // TODO не хранить в словаре а доставать пикчу
	charta, exists := (*s.db)[id]
	if !exists {
		return charta, errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	return charta, nil
}

func (s *Storage) Update(id uuid.UUID, newCharta models.Charta) error {
	charta, exists := (*s.db)[id]
	if !exists {
		return errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	newCharta.Id = charta.Id
	(*s.db)[id] = newCharta
	if err := s.WriteToFile(newCharta); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(id uuid.UUID) error {
	if _, exists := (*s.db)[id]; !exists {
		return errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	delete(*s.db, id)
	if err := os.Remove(fmt.Sprintf("%s.bmp", id.String())); err != nil {
		return err
	}
	return nil
}

func (s *Storage) WriteToFile(charta models.Charta) error {
	file, err := os.Create(fmt.Sprintf("%s.bmp", charta.Id))
	defer func(file *os.File) {
		if file.Close() != nil {
			return
		}
	}(file)
	if err != nil {
		return err
	}
	err = bmp.Encode(file, &charta.Image)
	if err != nil {
		return err
	}
	return nil
}
