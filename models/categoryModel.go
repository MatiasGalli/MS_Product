package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key"`
	Name string    `json:"name"`
}
