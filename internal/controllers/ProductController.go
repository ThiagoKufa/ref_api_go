package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ThiagoKufa/ref_api_go/internal/entities"
	service "github.com/ThiagoKufa/ref_api_go/internal/services"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := c.service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduct)
}

func (c *ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {

	products, err := c.service.GetALlProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
