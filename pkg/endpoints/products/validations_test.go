package productHandlers

import (
	// 	"main/pkg/models"
	"main/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulValidation(t *testing.T) {
	newProduct := models.Product{Name: "product1PR", Description: "This is the test product"}
	err := validateAddProduct(newProduct)
	assert.NoError(t, err)
}
