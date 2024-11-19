package middlewares

import (
	"fmt"
	"strings"

	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ApiRequestID(ctx *fiber.Ctx) error {
	/*-----------------------------------------------
	  01. GET VALUE OF THE 'Request-ID' HEADER
	 ------------------------------------------------*/
	var requestID string

	// First, check if 'event-hub-sign-auth' header is present
	requestID = ctx.Get("event-hub-sign-auth")
	if requestID == "" {
		// If 'event-hub-sign-auth' is not found, check 'Authorization' header
		authHeader := ctx.Get("Authorization")
		if strings.HasPrefix(authHeader, "event-hub-sign-auth: ") {
			// Extract the REQ-ID from the Authorization header
			requestID = strings.TrimPrefix(authHeader, "event-hub-sign-auth: ")
			requestID = strings.TrimSpace(requestID) // Remove any leading/trailing whitespace
		}
	}

	if requestID == "" {
		// If requestID is still empty, return unauthorized response
		return response.ErrorResponse("Missing Request ID", fiber.StatusUnauthorized, ctx)
	}

	utils.WarningPrint("REQ-ID: " + requestID)

	// 1. Print Request Headers
	headers := ctx.GetReqHeaders()
	fmt.Printf("Request Headers: %+v\n", headers)

	// 2. Print Raw Request Body
	rawBody := ctx.Body()
	fmt.Printf("Raw Request Body: %s\n", string(rawBody))

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
