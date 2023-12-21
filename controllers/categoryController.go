package controllers

import (
	db "github.com/MatiasGalli/MS_Product/config"
	"github.com/MatiasGalli/MS_Product/models"
)

func CreateCategory(category models.Category) (models.Category, error) {
	result := db.DB.Create(&category)
	return category, result.Error
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	result := db.DB.Find(&categories)
	return categories, result.Error
}

func GetCategory(categoryID int) (models.Category, error) {
	var category models.Category
	result := db.DB.Where(&category, "id = ?", categoryID)
	return category, result.Error
}

func UpdateCategory(categoryID int, category models.Category) (models.Category, error) {
	result := db.DB.Save(&category).Where("id = ?", categoryID)
	return category, result.Error
}

func DeleteCategory(categoryID int) error {
	var category models.Category
	result := db.DB.Delete(&category, "id = ?", categoryID)
	return result.Error
}
