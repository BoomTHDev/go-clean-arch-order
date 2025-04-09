package postgres

import (
	"boomth/internal/domain"

	"gorm.io/gorm"
)

type (
	OrderModel struct {
		Id    uint
		Total float64
	}

	orderRepository struct {
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (OrderModel) TableName() string {
	return "orders"
}

func (r *orderRepository) Save(order domain.Order) error {
	model := OrderModel{
		Total: order.Total,
	}

	result := r.db.Create(&model)
	return result.Error
}

func (r *orderRepository) FindAll() ([]domain.Order, error) {
	models := []OrderModel{}

	result := r.db.Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	orders := []domain.Order{}
	for _, model := range models {
		orders = append(orders, domain.Order{
			Id:    model.Id,
			Total: model.Total,
		})
	}

	return orders, nil
}
