package products

import (
	"main/logs"
	"main/pkg/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (productRepo *ProductRepository) CreateProduct(newProduct models.Product) {

	result := productRepo.Db.Create(&newProduct)
	if result.Error != nil {
		logs.ErrorLog("Error inserting product: %v , error %v", newProduct, result.Error)
	}
}

func (productRepo *ProductRepository) UpdateProduct(id string, updatedProduct *models.Product) {

	result := productRepo.Db.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct)
	if result.Error != nil {
		logs.ErrorLog("Error updating product: %v , error %v", updatedProduct, result.Error)
	}
}

func (productRepo *ProductRepository) GetProduct(id string) *models.Product {

	var product models.Product
	result := productRepo.Db.First(&product, id)
	if result.Error != nil {
		logs.ErrorLog("Error to get product with Id %v error: %v", id, result.Error)
		return nil
	}
	return &product
}

func (productRepo *ProductRepository) GetAllProduct() []models.Product {

	var products []models.Product

	result := productRepo.Db.Find(&products)
	if result.Error != nil {
		logs.ErrorLog("Error fetching all products, error: %v", result.Error)
		return nil
	}

	return products
}
