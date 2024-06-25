package services

import "main/pkg/models"

type ProductService interface {
	AddProduct(models.Product)
	GetProduct(string) models.Product
	GetAllProducts() []models.Product
	DeleteProduct(string)
}
