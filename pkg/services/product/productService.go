package services

import (
	"main/pkg/models"
	repository "main/pkg/repositories/products"
)

type productService struct {
	productRepository repository.Product
}

func NewProductService(repo repository.Product) ProductService {
	return &productService{
		productRepository: repo,
	}
}

func (s *productService) GetProduct(id string) models.Product {

	product := s.productRepository.GetProduct(id)
	return *product

}

func (s *productService) AddProduct(newProduct models.Product) {
	s.productRepository.CreateProduct(newProduct)
}

func (s *productService) UpdateProduct(id string, updatedProduct models.Product) {
	s.productRepository.UpdateProduct(id, updatedProduct)
}

func (s *productService) GetAllProducts() []models.Product {
	return s.productRepository.GetAllProducts()

}

func (s *productService) DeleteProduct(id string) {

}
