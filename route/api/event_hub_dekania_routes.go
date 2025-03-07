package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func NonAuthenticatedEventHubDekaniaRoutes(route fiber.Router) {
	route.Get("/get/all/regions", controllers.EventHubDekaniaController.GetAllRegions)
	route.Get("/get/all/dekania", controllers.EventHubDekaniaController.GetAllDekania)
}

func EventHubDekaniaRoutes(route fiber.Router) {
	route.Post("/add/region", controllers.EventHubDekaniaController.AddRegion)
	route.Post("/get/all/regions/by/pagination", controllers.EventHubDekaniaController.GetAllRegionsByPagination)
	route.Post("/update/region", controllers.EventHubDekaniaController.UpdateRegion)
	route.Post("/delete/region", controllers.EventHubDekaniaController.DeleteRegion)

	route.Post("/add/dekania", controllers.EventHubDekaniaController.AddDekania)
	route.Get("/get/all/dekania", controllers.EventHubDekaniaController.GetAllDekania)
	route.Post("/get/all/dekania/by/pagination", controllers.EventHubDekaniaController.GetAllDekaniaByPagination)
	route.Post("/update/dekania", controllers.EventHubDekaniaController.UpdateDekania)
	route.Post("/delete/dekania", controllers.EventHubDekaniaController.DeleteDekania)
}
