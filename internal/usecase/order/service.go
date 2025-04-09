package order

import (
	"boomth/internal/domain"
	"errors"
)

type (
	service struct {
		orderRepo domain.OrderRepository
		logger    Logger
	}

	Logger interface {
		Error(message any, fields ...any)
	}
)

func NewOrderUseCase(orderRepo domain.OrderRepository, logger Logger) UseCase {
	return &service{
		orderRepo: orderRepo,
		logger:    logger,
	}
}

func (s *service) CreateOrder(input CreateOrderInput) error {
	if input.Total <= 0 {
		s.logger.Error("Total must be positive", "total", input.Total)
		return errors.New("total must be positive")
	}

	order := domain.Order{
		Total: input.Total,
	}

	err := s.orderRepo.Save(order)
	if err != nil {
		s.logger.Error("Failed to create order", "error", err)
		return err
	}

	return nil
}

func (s *service) GetAllOrders() ([]OrderOutput, error) {
	orders, err := s.orderRepo.FindAll()
	if err != nil {
		s.logger.Error("Failed to fetch orders", "error", err)
		return nil, err
	}

	result := []OrderOutput{}
	for _, order := range orders {
		result = append(result, OrderOutput{
			Id:    order.Id,
			Total: order.Total,
		})
	}

	return result, nil
}
