package controllers

import (
	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/http/requests/payments"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubPaymentController = newEventHubPaymentController()

type eventHubPaymentController struct {
}

func newEventHubPaymentController() eventHubPaymentController {
	return eventHubPaymentController{}
}

func (c eventHubPaymentController) PushUSSD(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request payments.EventHubPaymentRequest
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

	return response.SuccessResponse("Configuration added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubPaymentController) GetPaymentTransactions(ctx *fiber.Ctx) error {
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
	 04. GET PAYMENT TRANSACTIONS AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	events, err := service.EventHubPaymentService.GetPaymentTransactions(pagination, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(events, fiber.StatusOK, ctx)
}
