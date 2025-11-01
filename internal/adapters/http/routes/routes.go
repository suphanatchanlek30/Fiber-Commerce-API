package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/adapters/http/handlers"
	"github.com/suphanatchanlek30/fiber-commerce-api/internal/adapters/http/middleware"
)

// Setup Route
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {

	// สร้าง handler สำหรับ admin
	adminHandler := handlers.NewAdminHandler()

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
	admin.Get("/dashboard", adminHandler.GetDashboard)
}
