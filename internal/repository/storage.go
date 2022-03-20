package repository

import (
	"errors"
	"fmt"
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"os"
	"path/filepath"
)

type Storage struct {
	directoryPath    string
	existingElements *map[uuid.UUID]bool
}

func NewStorage(path string) *Storage {
	return &Storage{directoryPath: path, existingElements: &map[uuid.UUID]bool{}}
}

func (s *Storage) Add(charta *models.Charta) (uuid.UUID, error) {
	id := uuid.New()
	charta.Id = id
	(*s.existingElements)[id] = true
	if err := s.WriteToFile(charta); err != nil {
		return [16]byte{}, err
	}
	return id, nil
}

func (s *Storage) GetById(id uuid.UUID) (*models.Charta, error) {
	var charta *models.Charta
	_, exists := (*s.existingElements)[id]
	if !exists {
		return charta, errors.New(fmt.Sprintf("ChartaRepository with id = %s doesn't exist", id))
	}
	charta, err := s.ReadFromFile(id.String())
	if err != nil {
		return charta, err
	}
	return charta, nil
}

func (s *Storage) Update(id uuid.UUID, newCharta *models.Charta) error {
	_, exists := (*s.existingElements)[id]
	if !exists {
		return errors.New(fmt.Sprintf("ChartaRepository with id = %s doesn't exist", id))
	}
	charta, err := s.ReadFromFile(id.String())
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
		return errors.New(fmt.Sprintf("ChartaRepository with id = %s doesn't exist", id))
	}
	delete(*s.existingElements, id)
	filePath := filepath.Join(s.directoryPath, fmt.Sprintf("%s.bmp", id.String()))
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func (s *Storage) ReadFromFile(fileName string) (*models.Charta, error) {
	var charta *models.Charta
	filePath := filepath.Join(s.directoryPath, fmt.Sprintf("%s.bmp", fileName))
	file, err := os.Open(filePath)
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
	id, err := uuid.Parse(fileName)
	if err != nil {
		return charta, err
	}
	charta = models.NewChartaWithId(id, decoded)
	return charta, nil
}

func (s *Storage) WriteToFile(charta *models.Charta) error {
	charta.Lock()
	defer charta.Unlock()
	filePath := filepath.Join(s.directoryPath, fmt.Sprintf("%s.bmp", charta.Id))
	file, err := os.Create(filePath)
	defer func(file *os.File) {
		if file.Close() != nil {
			return
		}
	}(file)
	if err != nil {
		return err
	}
	err = bmp.Encode(file, charta.Image)
	if err != nil {
		return err
	}
	return nil
}
