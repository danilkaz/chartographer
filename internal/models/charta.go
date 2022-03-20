package models

import (
	"github.com/danilkaz/chartographer/internal/models/errors"
	"github.com/google/uuid"
	"image"
	"math"
	"sync"
)

type Charta struct {
	Id    uuid.UUID
	Image *BitmapImage
	sync.Mutex
}

func NewCharta(img image.Image) *Charta {
	return &Charta{
		Image: NewBitmapImage(img),
	}
}

func NewChartaWithId(id uuid.UUID, img image.Image) *Charta {
	return &Charta{
		Id:    id,
		Image: NewBitmapImage(img),
	}
}

func (c *Charta) ChangePartOfImage(x, y, width, height int, otherCharta *Charta) error {
	bounds := c.Image.Bounds()
	startRow := int(math.Max(0, float64(y)))
	endRow := int(math.Min(float64(bounds.Dy()), float64(y+height)))
	startColumn := int(math.Max(0, float64(x)))
	endColumn := int(math.Min(float64(bounds.Dx()), float64(x+width)))
	isChanged := false
	for row := startRow; row < endRow; row++ {
		for column := startColumn; column < endColumn; column++ {
			c.Image.Set(column, row, otherCharta.Image.At(column-startColumn, row-startRow))
			isChanged = true
		}
	}
	if !isChanged {
		return errors.NotChangedError{}
	}
	return nil
}

func (c *Charta) GetSubImage(x, y, width, height int) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := c.Image.Bounds()
	if x+width <= 0 || y+height <= 0 || x >= bounds.Dx() || y >= bounds.Dy() {
		return img, errors.OutOfScopeError{}
	}
	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			img.Set(column, row, c.Image.At(column+x, row+y))
		}
	}
	return img, nil
}
