package repository

import (
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateOrder(order *model.Order) error {
	return r.db.Create(order).Error
}
