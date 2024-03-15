package controllers

import (
	"github.com/EventHubzTz/event_hub_service/app/http/requests/dekania"
	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubDekaniaController = newEventHubDekaniaController()

type eventHubDekaniaController struct {
}

func newEventHubDekaniaController() eventHubDekaniaController {
	return eventHubDekaniaController{}
}

func (c eventHubDekaniaController) AddDekania(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request dekania.EventHubDekaniaRequest
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
	 04. ADD DEKANIA
	-----------------------------------------------------------------------*/
	err = service.EventHubDekaniaService.AddDekania(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Dekania added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubDekaniaController) GetAllDekania(ctx *fiber.Ctx) error {

	dekania, dbErr := repositories.EventHubDekaniaRepository.GetAllDekania()

	if dbErr.RowsAffected == 0 {
		return response.ErrorResponse("No records found in dekania database", fiber.StatusOK, ctx)
	}

	return response.InternalServiceDataResponse(dekania, fiber.StatusOK, ctx)
}

func (c eventHubDekaniaController) GetAllDekaniaByPagination(ctx *fiber.Ctx) error {
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
	 04. GET DEKANIA AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dekania, err := service.EventHubDekaniaService.GetAllDekaniaByPagination(pagination, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(dekania, fiber.StatusOK, ctx)
}

func (c eventHubDekaniaController) UpdateDekania(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request dekania.EventHubDekaniaUpdateRequest
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
	 04. UPDATE DEKANIA NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err = service.EventHubDekaniaService.UpdateDekania(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Dekania updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubDekaniaController) DeleteDekania(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request dekania.EventHubDekaniaGetRequest
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
	 04. DELETE DEKANIA AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubDekaniaRepository.DeleteDekania(request.DekaniaID)
	/*---------------------------------------------------------
	 05. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete dekania", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Dekania deleted successfully!", fiber.StatusOK, ctx)
}
