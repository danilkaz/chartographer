package models

import "github.com/google/uuid"

type Charta struct {
	Id uuid.UUID
}

func NewCharta() *Charta {
	return &Charta{}
}
