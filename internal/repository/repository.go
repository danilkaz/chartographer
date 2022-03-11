package repository

import "github.com/danilkaz/chartographer/internal/models"

type Repository struct {
	models.ChartaInterface
}

func NewRepository(storage *Storage) *Repository { // TODO подумать как обобщить
	return &Repository{ChartaInterface: storage}
}
