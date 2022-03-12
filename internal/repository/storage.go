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
	mu               sync.Mutex
	existingElements *map[uuid.UUID]bool
}

func NewStorage(storage *map[uuid.UUID]bool) *Storage {
	return &Storage{existingElements: storage}
}

func (s *Storage) Add(charta models.Charta) (uuid.UUID, error) {
	id := uuid.New()
	charta.Id = id
	(*s.existingElements)[id] = true
	if err := s.WriteToFile(charta); err != nil {
		return [16]byte{}, err
	}
	return id, nil
}

func (s *Storage) GetById(id uuid.UUID) (models.Charta, error) {
	var charta models.Charta
	_, exists := (*s.existingElements)[id]
	if !exists {
		return charta, errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	charta, err := s.ReadFromFile(id)
	if err != nil {
		return charta, err
	}
	return charta, nil
}

func (s *Storage) Update(id uuid.UUID, newCharta models.Charta) error {
	var charta models.Charta
	_, exists := (*s.existingElements)[id]
	if !exists {
		return errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	charta, err := s.ReadFromFile(id)
	if err != nil {
		return err
	}
	newCharta.Id = charta.Id
	if err = s.WriteToFile(newCharta); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(id uuid.UUID) error {
	if _, exists := (*s.existingElements)[id]; !exists {
		return errors.New(fmt.Sprintf("Charta with id = %s doesn't exist", id))
	}
	delete(*s.existingElements, id)
	if err := os.Remove(fmt.Sprintf("%s.bmp", id.String())); err != nil {
		return err
	}
	return nil
}

func (s *Storage) ReadFromFile(id uuid.UUID) (models.Charta, error) {
	var charta models.Charta
	file, err := os.Open(fmt.Sprintf("%s.bmp", id))
	defer func(file *os.File) {
		if file.Close() != nil {
			return
		}
	}(file)
	if err != nil {
		return charta, err
	}
	decoded, err := bmp.Decode(file)
	if err != nil {
		return charta, err
	}
	return *models.NewCharta(decoded), nil
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
