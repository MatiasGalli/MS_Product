package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"not null;" json:"name"`
	Price       float32 `gorm:"not null;" json:"price"`
	Stock       int     `gorm:"not null;" json:"stock"`
	Description string  `gorm:"not null;" json:"description"`
	Offer       bool    `gorm:"not null;" json:"offer"`
	Promotion   float32 `gorm:"not null;" json:"promotion"`
	Image       string  `gorm:"not null;" json:"image"`
	CategoryID  int     `gorm:"not null;" json:"category_id"`

	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}
