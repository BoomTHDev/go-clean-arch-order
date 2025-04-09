package main

import (
	"boomth/internal/adapter/controller"
	"boomth/internal/adapter/repository/postgres"
	"boomth/internal/infrastructure/config"
	"boomth/internal/infrastructure/db"
	"boomth/internal/infrastructure/logging"
	"boomth/internal/usecase/order"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := setupApp(cfg)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}

func setupApp(cfg *config.Config) *fiber.App {
	// Infrastructure
	logger, err := logging.NewZapLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	database, err := db.NewPostgresDB(cfg.Database.DSN)
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
	orders.Post("/", orderController.CreateOrder)

	return app
}
