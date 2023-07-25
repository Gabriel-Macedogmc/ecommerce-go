package repository

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/products-backend/internal/domain"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetById(id uint) (*model.Product, error) {
	var product model.Product

	err := r.db.First(&product, id).Error

	if err != nil {
		return nil, &domain.ErrNotFound{EntityType: "Product", EntityID: fmt.Sprintf("%d", id)}
	}

	return &product, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(product *model.Product) error {
	return r.db.Delete(product).Error
}
