package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubEventsManagementRoutes(route fiber.Router) {
	route.Post("/add/event", controllers.EventHubEventsManagementController.AddEvent)
	route.Post("/add/event/image", controllers.EventHubEventsManagementController.AddEventImage)
	route.Post("/get/events", controllers.EventHubEventsManagementController.GetEvents)
	route.Post("/get/event", controllers.EventHubEventsManagementController.GetEvent)
	route.Post("/update/event", controllers.EventHubEventsManagementController.UpdateEvent)
	route.Post("/delete/event/image", controllers.EventHubEventsManagementController.DeleteEventImage)
	route.Post("/delete/event", controllers.EventHubEventsManagementController.DeleteEvent)

	route.Post("/add/event/package", controllers.EventHubEventsManagementController.AddEventPackage)
	route.Post("/get/event/packages", controllers.EventHubEventsManagementController.GetAllEventPackages)
	route.Post("/update/event/package", controllers.EventHubEventsManagementController.UpdateEventPackage)
	route.Post("/delete/event/package", controllers.EventHubEventsManagementController.DeleteEventPackage)

	route.Post("/get/dashboard/statistics", controllers.EventHubEventsManagementController.GetDashboardStatistics)
}
