package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	CategoryID  string    `json:"category_id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Offer       bool      `json:"offer"`
	Promotion   float32   `json:"promotion"`
	Image       string    `json:"image"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}
