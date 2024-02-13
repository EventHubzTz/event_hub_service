package provider

import (
	"github.com/EventHubzTz/event_hub_service/app/http/middlewares"
	"github.com/EventHubzTz/event_hub_service/route/api"
	"github.com/gofiber/fiber/v2"
)

func RouteProvider(app *fiber.App) {

	apiRoute := app.Group("/api/v1", middlewares.Api)
	api.ApiRoute(apiRoute)
	api.EventHubRequestIDRoutes(apiRoute)
	apiExternalOperationRoute := app.Group("/api/v1", middlewares.ApiRequestID)
	api.EventHubUsersManagementRoutes(apiExternalOperationRoute)
	apiAuthenticatedRoute := app.Group("/api", middlewares.ApiAuth)
	api.AuthenticatedEventHubUsersManagementRoutes(apiAuthenticatedRoute)
}
