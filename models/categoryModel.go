package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `gorm:"primary_key" json:"id"`
	Name string    `gorm:"not null;" json:"name"`
}
