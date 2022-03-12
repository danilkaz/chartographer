package models

import (
	"github.com/google/uuid"
	"image"
	"math"
)

type Charta struct {
	Id    uuid.UUID
	Image BitmapImage
}

func NewCharta(img image.Image) *Charta {
	return &Charta{
		Image: *NewBitmapImage(img),
	}
}

func (c *Charta) ChangePartOfImage(x, y, width, height int, otherCharta Charta) {
	bounds := c.Image.Bounds()
	startRow := int(math.Max(0, float64(y)))
	endRow := int(math.Min(float64(bounds.Dy()), float64(y+height)))
	startColumn := int(math.Max(0, float64(x)))
	endColumn := int(math.Min(float64(bounds.Dx()), float64(x+width)))
	for row := startRow; row < endRow; row++ {
		for column := startColumn; column < endColumn; column++ {
			c.Image.Set(column, row, otherCharta.Image.At(column-startColumn, row-startRow))
		}
	}
}

func (c *Charta) SubCharta(x, y, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			img.Set(column, row, c.Image.At(column+x, row+y))
		}
	}
	return img
}
