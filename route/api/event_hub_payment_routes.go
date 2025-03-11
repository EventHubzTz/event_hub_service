package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func NonAuthenticatedEventHubPaymentRoutes(route fiber.Router) {
	route.Post("/Checkout/Callback", controllers.EventHubPaymentController.UpdatePaymentStatus)
}

func EventHubPaymentRoutes(route fiber.Router) {
	route.Post("/push/ussd", controllers.EventHubPaymentController.VotingPushUSSD)
	route.Post("/get/voting/payment/transactions", controllers.EventHubPaymentController.GetVotingPaymentTransactions)
}

func AuthenticatedEventHubPaymentRoutes(route fiber.Router) {
	route.Post("/add/debit", controllers.EventHubPaymentController.AddDebit)
	route.Post("/pugu/marathon/push/ussd", controllers.EventHubPaymentController.PushUSSD)
	route.Post("/pugu/marathon/contribution/push/ussd", controllers.EventHubPaymentController.ContributionPushUSSD)
	route.Post("/get/payment/transaction", controllers.EventHubPaymentController.GetPaymentTransactions)
	route.Post("/get/contribution/transactions", controllers.EventHubPaymentController.GetContributionTransactions)
	route.Post("/get/transaction/by/transaction/id", controllers.EventHubPaymentController.GetTransactionByTransactionID)
	route.Post("/get/contribution/by/transaction/id", controllers.EventHubPaymentController.GetContributionByTransactionID)
}
