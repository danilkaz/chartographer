package models

import (
	"github.com/google/uuid"
	"image"
	"image/color"
)

type Charta struct {
	Id            uuid.UUID
	Image         image.Image
	ChangedPixels map[image.Point]color.Color
}

func NewCharta(img image.Image) *Charta {
	return &Charta{
		Image:         img,
		ChangedPixels: map[image.Point]color.Color{},
	}
}
