package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func ApiRoute(route fiber.Router) {
	route.Group("")
	route.Get("/", controllers.Index)
}
