package mock

import "boomth/internal/domain"

type (
	orderRepository struct {
		orders []domain.Order
	}
)

func NewOrderRepository() domain.OrderRepository {
	return &orderRepository{
		orders: []domain.Order{
			{Id: 1, Total: 100},
			{Id: 2, Total: 200},
		},
	}
}

func (r *orderRepository) Save(order domain.Order) error {
	newOrder := domain.Order{
		Id:    uint(len(r.orders) + 1),
		Total: order.Total,
	}

	r.orders = append(r.orders, newOrder)
	return nil
}

func (r *orderRepository) FindAll() ([]domain.Order, error) {
	return r.orders, nil
}
