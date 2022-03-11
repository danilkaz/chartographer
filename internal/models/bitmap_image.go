package models

import (
	"image"
	"image/color"
)

type BitmapImage struct {
	image.Image
	changedPixels map[image.Point]color.Color
}

func NewBitmapImage(img image.Image) *BitmapImage {
	return &BitmapImage{img, map[image.Point]color.Color{}}
}

func (b *BitmapImage) Set(x, y int, clr color.Color) {
	b.changedPixels[image.Point{X: x, Y: y}] = clr
}

func (b *BitmapImage) At(x, y int) color.Color {
	if clr, exists := b.changedPixels[image.Point{X: x, Y: y}]; exists {
		return clr
	}
	return b.Image.At(x, y)
}
