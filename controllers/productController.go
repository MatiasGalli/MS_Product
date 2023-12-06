package controllers

import (
	db "github.com/MatiasGalli/MS_Product/config"
	"github.com/MatiasGalli/MS_Product/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	result := db.DB.Create(&product)
	return product, result.Error
}
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := db.DB.Find(&products)
	return products, result.Error
}

func GetProduct(productID string) (models.Product, error) {
	var product models.Product
	result := db.DB.Where(&product, "id = ?", productID)
	return product, result.Error
}

func UpdateProduct(productID, product models.Product) (models.Product, error) {
	result := db.DB.Save(&product).Where("id = ?", productID)
	return product, result.Error
}

func DeleteProduct(productID string) error {
	var product models.Product
	result := db.DB.Delete(&product, "id = ?", productID)
	return result.Error
}
