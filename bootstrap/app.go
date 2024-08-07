package bootstrap

import (
	"os"

	"github.com/EventHubzTz/event_hub_service/app/provider"
	"github.com/EventHubzTz/event_hub_service/database"
	"github.com/EventHubzTz/event_hub_service/database/migrations"
	// "github.com/EventHubzTz/event_hub_service/database/seeders"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func StartApp() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file is missing")
	}

	app := fiber.New(fiber.Config{BodyLimit: 1024 * 1024 * 1024})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/", "./public")

	utils.SuccessPrint("We about to connect Database")
	database.DatabaseConnection()
	migrations.Migrate()
	// seeders.Seed()
	defer database.CloseDB()

	repositories.Init()

	provider.RouteProvider(app)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
