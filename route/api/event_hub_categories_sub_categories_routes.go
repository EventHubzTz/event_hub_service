package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubCategoriesSubCategoriesRoutes(route fiber.Router) {
	route.Post("/create/event/category", controllers.EventHubCategoriesSubCategoriesController.CreateEventCategory)
	route.Get("/get/all/event/categories", controllers.EventHubCategoriesSubCategoriesController.GetAllEventCategories)
	route.Post("/get/all/event/categories/by/pagination", controllers.EventHubCategoriesSubCategoriesController.GetAllEventCategoriesByPagination)
	route.Post("/update/event/category", controllers.EventHubCategoriesSubCategoriesController.UpdateEventCategory)
	route.Post("/delete/event/category", controllers.EventHubCategoriesSubCategoriesController.DeleteEventCategory)

	route.Post("/create/event/subcategory", controllers.EventHubCategoriesSubCategoriesController.CreateEventSubCategory)
	route.Post("/get/all/event/subcategories", controllers.EventHubCategoriesSubCategoriesController.GetAllEventSubCategories)
	route.Post("/get/all/event/subcategories/by/pagination", controllers.EventHubCategoriesSubCategoriesController.GetAllEventSubCategoriesByPagination)
	route.Post("/update/event/subcategory", controllers.EventHubCategoriesSubCategoriesController.UpdateEventSubCategory)
	route.Post("/delete/event/subcategory", controllers.EventHubCategoriesSubCategoriesController.DeleteEventSubCategory)
}
