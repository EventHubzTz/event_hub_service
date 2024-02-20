package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubConfigurationsRoutes(route fiber.Router) {
	route.Post("/add/configuration", controllers.EventHubConfigurationsController.AddConfiguration)
	route.Get("/get/configurations", controllers.EventHubConfigurationsController.GetConfigurations)
}
