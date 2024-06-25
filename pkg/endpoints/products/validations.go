package productHandlers

import (
	"errors"
	"main/pkg/models"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

func validateAddProduct(productToValidate models.Product) error {

	err := validation.Validate(productToValidate.Name, validation.Required, mustHavePRsuffixedRule())
	return err
}

type mustHavePRsuffixedCheck struct{}

func (mustHavePRsuffixedCheck) Validate(productName interface{}) error {
	productNameString := productName.(string)
	if !strings.Contains(productNameString, "PR") {
		return errors.New("product should have PR suffix")
	}
	return nil

}

func mustHavePRsuffixedRule() validation.Rule {
	return mustHavePRsuffixedCheck{}
}
