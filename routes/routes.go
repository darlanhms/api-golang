package routes

import "github.com/gofiber/fiber/v2"

// Register all routes of the App
func Register(app *fiber.App) {
	RegisterProducts(app)
}
