package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/naphat/fiber-ecommerce-api/internal/adapters/http/middleware"
)

// Setup Route
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API routes
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.Get("/profile", authHandler.GetProfile)

	// Admin only routes
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireRole("admin"))
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the admin dashboard",
		})
	})
}