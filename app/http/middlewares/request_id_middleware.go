package middlewares

import (
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ApiRequestID(ctx *fiber.Ctx) error {
	/*-----------------------------------------------
	 01. GET VALUE OF THE 'Request-ID' HEADER
	------------------------------------------------*/
	requestID := ctx.Get("event-hub-sign-auth")
	utils.WarningPrint("REQ-ID: " + requestID)
	/*-------------------------------------------------
	02. COMPARE TO THE ACTIVE REQUEST ID IN SYSTEM
	--------------------------------------------------*/
	status, err := service.EventHubRequestIDService.VerifyRequestID(requestID)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusUnauthorized, ctx)
	}
	if !status {
		return response.ErrorResponse("Unauthorized Access", fiber.StatusUnauthorized, ctx)
	}
	return ctx.Next()
}
