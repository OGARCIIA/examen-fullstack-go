package application

import (
	"github.com/OGARCIIA/examen-backend/domain"
	"github.com/OGARCIIA/examen-backend/infrastructure"
)

func GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := infrastructure.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *domain.Product) error {
	result := infrastructure.DB.Create(product)
	return result.Error
}

func UpdateProduct(id uint, updated *domain.Product) error {
	var product domain.Product
	result := infrastructure.DB.First(&product, id)
	if result.Error != nil {
		return result.Error
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price
	product.Stock = updated.Stock

	result = infrastructure.DB.Save(&product)
	return result.Error
}

func DeleteProduct(id uint) error {
	result := infrastructure.DB.Delete(&domain.Product{}, id)
	return result.Error
}
