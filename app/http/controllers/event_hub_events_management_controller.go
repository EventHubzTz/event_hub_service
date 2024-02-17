package controllers

import (
	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubEventsManagementController = newEventHubEventsManagementController()

type eventHubEventsManagementController struct {
}

func newEventHubEventsManagementController() eventHubEventsManagementController {
	return eventHubEventsManagementController{}
}

func (c eventHubEventsManagementController) AddEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
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
	 04. ADD EVENT
	-----------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.AddEvent(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Event added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetEvents(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventsGetsRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	var pagination models.Pagination
	pagination.Limit = request.Limit
	pagination.Sort = request.Sort
	pagination.Page = request.Page

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. GET EVENTS AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	events, err := service.EventHubEventsManagementService.GetEvents(pagination, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(events, fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetEvent(ctx *fiber.Ctx) error {
	/*--------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST UPDATING PASSWORD
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*---------------------------------------------------------
	 04. GET THE EVENT FROM THE DATABASE USING EVENT ID
	----------------------------------------------------------*/
	event := service.EventHubEventsManagementService.GetEvent(request.EventID)
	if event == nil {
		return response.ErrorResponse("Event details not found in the system", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 09. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.InternalServiceDataResponse(event, fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) UpdateEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubUpdateEventRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. UPDATE EVENT NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.UpdateEvent(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) DeleteEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. DELETE EVENT AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubEventsManagementRepository.DeleteEvent(request.EventID)
	/*---------------------------------------------------------
	 05. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete event", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event deleted successfully!", fiber.StatusOK, ctx)
}
