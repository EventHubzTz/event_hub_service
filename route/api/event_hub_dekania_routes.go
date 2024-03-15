package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubDekaniaRoutes(route fiber.Router) {
	route.Post("/add/dekania", controllers.EventHubDekaniaController.AddDekania)
	route.Get("/get/all/dekania", controllers.EventHubDekaniaController.GetAllDekania)
	route.Post("/get/all/dekania/by/pagination", controllers.EventHubDekaniaController.GetAllDekaniaByPagination)
	route.Post("/update/dekania", controllers.EventHubDekaniaController.UpdateDekania)
	route.Post("/delete/dekania", controllers.EventHubDekaniaController.DeleteDekania)
}
