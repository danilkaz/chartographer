package models

import "github.com/google/uuid"

type ChartaInterface interface {
	Create(width, height int) (uuid.UUID, error)
	SaveRestoredFragment(x, y, width, height int) error
	GetPart(x, y, width, height int) (Charta, error)
	Delete() error
}

type Charta struct {
	Id uuid.UUID
}

func NewCharta() *Charta {
	return &Charta{}
}
