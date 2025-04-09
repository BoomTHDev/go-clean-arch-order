package domain

type (
	Order struct {
		Id    uint    `json:"id"`
		Total float64 `json:"total"`
	}

	OrderRepository interface {
		Save(order Order) error
		FindAll() ([]Order, error)
	}
)
