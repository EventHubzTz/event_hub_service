package controllers

import (
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/gofiber/fiber/v2"
)

var EventHubRequestIDController = newEventHubRequestIDManagementController()

type eventHubRequestIDController struct {
}

func newEventHubRequestIDManagementController() eventHubRequestIDController {
	return eventHubRequestIDController{}
}

func (s eventHubRequestIDController) GetRequestID(ctx *fiber.Ctx) error {
	/*--------------------------------------------------------
	 01. RETURNING THE RESPONSE WITH MICRO SERVICE REQUEST ID
	----------x------------------------------------------------*/
	return response.MapDataResponse(service.EventHubRequestIDService.GetRequestID(), fiber.StatusOK, ctx)
}
