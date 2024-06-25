package repository

import "main/pkg/models"

type Product interface {
	CreateProduct(newProduct models.Product)
	GetProduct(id string) *models.Product
	UpdateProduct(id string, updatedProduct models.Product)
	DeleteProduct(id string)
	GetAllProducts() []models.Product
}
