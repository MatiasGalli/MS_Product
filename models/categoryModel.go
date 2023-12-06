package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	ID   uuid.UUID `json:"id" gorm:"primary_key"`
	Name string    `json:"name"`
}
