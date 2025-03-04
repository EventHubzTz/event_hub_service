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
	api.NonAuthenticatedEventHubPaymentRoutes(apiRoute)
	apiExternalOperationRoute := app.Group("/api/v1", middlewares.ApiRequestID)
	api.EventHubUsersManagementRoutes(apiExternalOperationRoute)
	api.EventHubPaymentRoutes(apiExternalOperationRoute)
	api.NonAuthenticatedEventHubDekaniaRoutes(apiExternalOperationRoute)
	apiAuthenticatedRoute := app.Group("/api/v1", middlewares.ApiAuth)
	api.AuthenticatedEventHubUsersManagementRoutes(apiAuthenticatedRoute)
	api.EventHubEventsManagementRoutes(apiAuthenticatedRoute)
	api.EventHubCategoriesSubCategoriesRoutes(apiAuthenticatedRoute)
	api.AuthenticatedEventHubPaymentRoutes(apiAuthenticatedRoute)
	api.EventHubConfigurationsRoutes(apiAuthenticatedRoute)
	api.EventHubDekaniaRoutes(apiAuthenticatedRoute)
}
