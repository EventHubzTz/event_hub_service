package controllers

import (
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	return response.SuccessResponse("Service is online", fiber.StatusOK, ctx)
}
