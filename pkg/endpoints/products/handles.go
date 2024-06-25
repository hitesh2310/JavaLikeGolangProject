package productHandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"main/pkg/models"
	services "main/pkg/services/product"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service services.ProductService
}

func NewHandler(service services.ProductService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body: ", err)
		return
	}

	var newProduct models.Product

	if err := json.Unmarshal(body, &newProduct); err != nil {
		fmt.Println("Error unmarshalling into product obj: ", err)
		return
	}

	validateAddProduct(newProduct)
	h.service.AddProduct(newProduct)

}
