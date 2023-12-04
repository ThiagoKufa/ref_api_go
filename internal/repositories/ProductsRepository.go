package products

import (
	"github.com/ThiagoKufa/ref_api_go/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) CreateProduct(product *entities.Product) (*entities.Product, error) {
	result := repo.db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetAllProducts() ([]*entities.Product, error) {
	var products []*entities.Product

	result := repo.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
