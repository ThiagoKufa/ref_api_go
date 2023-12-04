package services

import (
	"github.com/ThiagoKufa/ref_api_go/internal/entities"
)

type ProductRepository interface {
	CreateProduct(product *entities.Product) (*entities.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *entities.Product) (*entities.Product, error) {
	return s.repo.CreateProduct(product)
}
