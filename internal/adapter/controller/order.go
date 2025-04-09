package controller

import (
	"boomth/internal/domain"
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

	switch err.(type) {
	case domain.ValidationError:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	case nil:
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Order created successfully",
		})
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
}

func (c *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := c.orderUseCase.GetAllOrders()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse("Failed to fetch orders"))
	}

	return ctx.Status(fiber.StatusOK).JSON(SuccessResponse(orders))
}
