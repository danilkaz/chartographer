package service

import (
	"github.com/danilkaz/chartographer/internal/models"
	"github.com/danilkaz/chartographer/internal/repository"
	"github.com/google/uuid"
	"image"
)

type ChartaService struct {
	repository.ChartaRepository
}

func NewChartaService(repo *repository.Repository) *ChartaService {
	return &ChartaService{repo}
}

func (s *ChartaService) Create(width, height int) (uuid.UUID, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	charta := models.NewCharta(img)
	return s.ChartaRepository.Add(charta)
}

func (s *ChartaService) SaveRestoredFragment(id uuid.UUID, x, y, width, height int, fragment *models.Charta) error {
	charta, err := s.ChartaRepository.GetById(id)
	if err != nil {
		return err
	}
	err = charta.ChangePartOfImage(x, y, width, height, fragment)
	if err != nil {
		return err
	}
	if err = s.ChartaRepository.Update(id, charta); err != nil {
		return err
	}
	return nil
}

func (s *ChartaService) GetPart(id uuid.UUID, x, y, width, height int) (*models.Charta, error) {
	charta, err := s.ChartaRepository.GetById(id)
	if err != nil {
		return charta, err
	}
	subImage, err := charta.GetSubImage(x, y, width, height)
	subImageCharta := &models.Charta{Image: models.NewBitmapImage(subImage)}
	if err != nil {
		return subImageCharta, err
	}
	return subImageCharta, nil
}

func (s *ChartaService) Delete(id uuid.UUID) error {
	return s.ChartaRepository.Delete(id)
}
