package main

import (
	"boomth/internal/adapter/controller"
	"boomth/internal/adapter/repository/postgres"
	"boomth/internal/infrastructure/db"
	"boomth/internal/infrastructure/logging"
	"boomth/internal/usecase/order"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Infrastructure
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	dsn := "host=localhost user=postgres password=8231 dbname=gorm_boom port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := db.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Repository
	orderRepo := postgres.NewOrderRepository(database)

	// Use Cases
	orderUseCase := order.NewOrderUseCase(orderRepo, logger)

	// Controller
	orderController := controller.NewOrderController(orderUseCase)

	// HTTP Server
	app := fiber.New()
	v1 := app.Group("/v1")

	orders := v1.Group("/orders")
	orders.Get("/", orderController.GetAllOrders)

	log.Fatal(app.Listen(":3001"))
}
