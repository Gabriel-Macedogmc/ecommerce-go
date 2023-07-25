package services

import (
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/model"
	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository}
}

func (s *OrderService) CreateOrder(data model.Order) error {
	err := s.orderRepository.CreateOrder(&data)
	if err != nil {
		return err
	}
	return nil
}
