package service

import (
	"github.com/google/uuid"
	"image"
	"image/color"
)

type BitmapImage struct {
	uuid.UUID
	image.Image
	colors map[image.Point]color.Color
}

func NewBitmapImage(bitmapImage image.Image) *BitmapImage {
	return &BitmapImage{uuid.New(), bitmapImage, map[image.Point]color.Color{}}
}

func (b *BitmapImage) SetColor(x, y int, color color.Color) {
	b.colors[image.Point{X: x, Y: y}] = color
}

func (b *BitmapImage) At(x, y int) color.Color {
	if c := b.colors[image.Point{X: x, Y: y}]; c != nil {
		return c
	}
	return b.Image.At(x, y)
}

func (b *BitmapImage) ChangePartOfImage(x, y, width, height int, otherImage BitmapImage) {
	for row := y; row < y+height; row++ {
		for column := x; column < x+width; column++ {
			b.SetColor(row, column, otherImage.At(row, column))
		}
	}
}
