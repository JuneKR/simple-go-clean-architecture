package usecases

import "neptune/entities"

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

// Create instance of OrderService
func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}

// CreateOrder implements OrderUseCase
func (s *OrderService) CreateOrder(order entities.Order) error {
	return s.repo.Save(order)
}
