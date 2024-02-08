package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubUsersManagementRoutes(route fiber.Router) {
	route.Post("/register/user", controllers.EventHubUsersManagementController.RegisterUser)
	route.Post("/login/user", controllers.EventHubUsersManagementController.LoginUser)
	route.Post("/get/users", controllers.EventHubUsersManagementController.GetUsers)
	route.Post("/change/password", controllers.EventHubUsersManagementController.ChangePassword)
}
