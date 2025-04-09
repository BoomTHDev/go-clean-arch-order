package controller

import (
	"boomth/internal/usecase/order"

	"github.com/gofiber/fiber/v2"
)

type (
	OrderController struct {
		orderUseCase order.UseCase
	}
)

func NewOrderController(useCase order.UseCase) *OrderController {
	return &OrderController{
		orderUseCase: useCase,
	}
}

func (c *OrderController) CreateOrder(ctx *fiber.Ctx) error {
	input := order.CreateOrderInput{}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	err := c.orderUseCase.CreateOrder(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order created successfully",
	})
}

func (c *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := c.orderUseCase.GetAllOrders()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch orders",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(orders)
}
