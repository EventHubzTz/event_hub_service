package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubRequestIDRoutes(route fiber.Router) {
	route.Get("/get/request/id", controllers.EventHubRequestIDController.GetRequestID)
}
