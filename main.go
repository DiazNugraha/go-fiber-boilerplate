package main

import (
	"go-fiber-boilerplate/pkg/configs"
	"go-fiber-boilerplate/pkg/middleware"
	"go-fiber-boilerplate/pkg/routes"
	"go-fiber-boilerplate/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
