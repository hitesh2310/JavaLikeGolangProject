package router

import (
	"log"
	"main/pkg/database"
	productHandlers "main/pkg/endpoints/products"
	repository "main/pkg/repositories/products"
	services "main/pkg/services/product"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return &http.Server{Addr: "localhost:" + port, Handler: newHandler()}
}

func newHandler() *httprouter.Router {
	r := httprouter.New()
	addRoutes(r)

	return r
}

func addRoutes(r *httprouter.Router) {
	productDbConn := database.GetDBConn()
	productRepository := repository.NewProductRepository(productDbConn)
	productService := services.NewProductService(productRepository)
	productHandlers := productHandlers.NewHandler(productService)
	r.POST("/add", productHandlers.AddProduct)

}
