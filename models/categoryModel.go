package models

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;" json:"name"`
}

func (Category) TableName() string {
	return "category"
}
