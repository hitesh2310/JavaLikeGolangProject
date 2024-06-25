package services

import (
	"fmt"
	"main/pkg/database"
	"main/pkg/models"
	repository "main/pkg/repositories/products"

	"testing"

	"github.com/bmizerany/assert"
)

// func init() {
// 	fmt.Println("Starting ECom App....")
// 	config.SetUpApplication()
// }

func TestInsertProduct(t *testing.T) {
	dbConn := database.GetDBConnQA()
	productRepo := repository.NewProductRepository(dbConn)
	productService := NewProductService(productRepo)
	fmt.Println(" Conn , Repo and Service is initialised.")

	newProduct := models.Product{Name: "product1", Description: "This is the second product"}
	productService.AddProduct(newProduct)

}

func TestGetAllProducts(t *testing.T) {
	dbConn := database.GetDBConnQA()
	productRepo := repository.NewProductRepository(dbConn)
	productService := NewProductService(productRepo)

	productList := productService.GetAllProducts()

	fmt.Println("ProductList", productList)

	for index, val := range productList {
		fmt.Println("i ", index, "product :", val)
	}
	assert.Equal(t, len(productList) > 0, true)
}
