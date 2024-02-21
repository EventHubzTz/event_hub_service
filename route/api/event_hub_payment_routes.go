package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubPaymentRoutes(route fiber.Router) {
	route.Post("/push/ussd", controllers.EventHubPaymentController.PushUSSD)
	route.Post("/get/payment/transaction", controllers.EventHubPaymentController.GetPaymentTransactions)
	route.Post("/Checkout/Callback", controllers.EventHubPaymentController.UpdatePaymentStatus)
}
