package order

type (
	CreateOrderInput struct {
		Total float64 `json:"total"`
	}

	OrderOutput struct {
		Id    uint    `json:"id"`
		Total float64 `json:"total"`
	}

	UseCase interface {
		CreateOrder(CreateOrderInput) error
		GetAllOrders() ([]OrderOutput, error)
	}
)
