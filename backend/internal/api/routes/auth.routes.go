package routes

import (
	"github.com/bibhu20031/CollabPad/backend/internal/api/controllers"
	middleware "github.com/bibhu20031/CollabPad/backend/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Public routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to CollabPad!") // Simple welcome message
	})
	// Protected routes
	app.Use("/private", middleware.Protected())

	app.Get("/private/profile", func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{"user": user})
	})
}
