package controllers

import (
	"github.com/EventHubzTz/event_hub_service/app/http/requests/configurations"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubConfigurationsController = newEventHubConfigurationsController()

type eventHubConfigurationsController struct {
}

func newEventHubConfigurationsController() eventHubConfigurationsController {
	return eventHubConfigurationsController{}
}

func (c eventHubConfigurationsController) AddConfiguration(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request configurations.EventHubConfigurationsRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*--------------------------------------------------------------------
	 04. ADD CONFIGURATION
	-----------------------------------------------------------------------*/
	err = service.EventHubConfigurationsService.AddConfiguration(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Configuration added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubConfigurationsController) GetConfigurations(ctx *fiber.Ctx) error {
	/*---------------------------------------------------------
	 01. GET CONFIGURATIONS
	----------------------------------------------------------*/
	configurations := service.EventHubConfigurationsService.GetConfigurations()
	if configurations == nil {
		return response.ErrorResponseStr("No records found !", fiber.StatusBadRequest, ctx)
	}
	return response.MapDataResponse(configurations, fiber.StatusOK, ctx)
}
