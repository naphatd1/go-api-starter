// @title Fiber Auth API
// @version 1.0
// @description Authentication API with Role-based Access Control
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"log"

	_ "github.com/naphat/fiber-ecommerce-api/docs"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/http/routes"
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/persistence/repositories"
	"github.com/naphat/fiber-ecommerce-api/internal/config"
	"github.com/naphat/fiber-ecommerce-api/internal/core/services"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// Setup database connection
	db := config.SetupDatabase(cfg)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app, authHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}