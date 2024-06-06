package routes

import (
	"go-fiber-boilerplate/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/generate", controllers.GeneratePages)
}
