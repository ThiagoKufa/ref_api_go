package main

import (
	"net/http"

	controller "github.com/ThiagoKufa/ref_api_go/internal/controllers"
	"github.com/ThiagoKufa/ref_api_go/internal/infra/database"
	repository "github.com/ThiagoKufa/ref_api_go/internal/repositories"
	service "github.com/ThiagoKufa/ref_api_go/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := database.ConectDatabase()
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productController.CreateProduct)
		r.Get("/", productController.GetAllProduct)

	})

	http.ListenAndServe(":3000", r)
}
