package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubEventsManagementRoutes(route fiber.Router) {
	route.Post("/add/event", controllers.EventHubEventsManagementController.AddEvent)
	route.Post("/get/events", controllers.EventHubEventsManagementController.GetEvents)
	route.Post("/get/event", controllers.EventHubEventsManagementController.GetEvent)
	route.Get("/update/event", controllers.EventHubEventsManagementController.UpdateEvent)
	route.Post("/delete/event", controllers.EventHubEventsManagementController.DeleteEvent)
}
