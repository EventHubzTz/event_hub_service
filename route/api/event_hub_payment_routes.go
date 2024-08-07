package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func NonAuthenticatedEventHubPaymentRoutes(route fiber.Router) {
	route.Post("/Checkout/Callback", controllers.EventHubPaymentController.UpdatePaymentStatus)
}

func EventHubPaymentRoutes(route fiber.Router) {
	route.Post("/make/vote", controllers.EventHubPaymentController.VotingPushUSSD)
}

func AuthenticatedEventHubPaymentRoutes(route fiber.Router) {
	route.Post("/push/ussd", controllers.EventHubPaymentController.PushUSSD)
	route.Post("/get/payment/transaction", controllers.EventHubPaymentController.GetPaymentTransactions)
	route.Post("/get/transaction/by/transaction/id", controllers.EventHubPaymentController.GetTransactionByTransactionID)
}
