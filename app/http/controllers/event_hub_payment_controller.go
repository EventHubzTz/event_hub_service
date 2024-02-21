package controllers

import (
	"strconv"
	"time"

	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/http/requests/payments"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils"
	"github.com/EventHubzTz/event_hub_service/utils/constants"
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
	userFromLocal := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	request.UserID = userFromLocal.Id
	request.Currency = constants.Currency
	request.OrderID = utils.GenerateOrderId()
	request.Provider = utils.CheckMobileNetwork(request.PhoneNumber)
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*----------------------------------------
	 04. CHECK IF EVENT EXIST IN THE SYSTEM
	------------------------------------------*/
	event, eventError := service.EventHubEventsManagementService.GetEvent(request.EventID)
	if eventError != nil {
		return response.ErrorResponse(eventError.Error(), fiber.StatusBadRequest, ctx)
	}
	request.Amount = event.EventEntrance
	/*---------------------------------------------------------
	 04. GET CONFIGURATIONS
	----------------------------------------------------------*/
	configurations := service.EventHubConfigurationsService.GetConfigurations()
	if configurations == nil {
		return response.ErrorResponseStr("No records found !", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 05. GET AZAMPAY API KEY
	----------------------------------------------------------*/
	apiKey, apiKeyError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(10)
	if apiKeyError != nil {
		return response.ErrorResponseStr(apiKeyError.Error.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 06. CHECK IF TOKEN TIME HAS EXPIRED
	----------------------------------------------------------*/
	generatedTime, err := time.Parse(time.RFC3339, configurations.AzampayTokenGeneratedTime)
	if err != nil {
		return response.ErrorResponseStr(err.Error(), fiber.StatusBadRequest, ctx)
	}
	if time.Now().After(generatedTime) {
		/*---------------------------------------------------------
		 07. GET AZAMPAY AUTHENTICATOR BASE URL
		----------------------------------------------------------*/
		authenticatorBaseURL, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(5)
		if err != nil {
			return response.ErrorResponseStr(err.Error.Error(), fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 08. GET AZAMPAY APP NAME
		----------------------------------------------------------*/
		appName, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(7)
		if err != nil {
			return response.ErrorResponseStr(err.Error.Error(), fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 09. GET AZAMPAY CLIENT ID
		----------------------------------------------------------*/
		clientID, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(8)
		if err != nil {
			return response.ErrorResponseStr(err.Error.Error(), fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 10. GET AZAMPAY CLIENT SECRET
		----------------------------------------------------------*/
		clientSecret, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(9)
		if err != nil {
			return response.ErrorResponseStr(err.Error.Error(), fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 11. GET BEARER TOKEN
		----------------------------------------------------------*/
		url := authenticatorBaseURL + "/AppRegistration/GenerateToken"
		tokenResponse, tokenError := helpers.GenerateAzamPayToken(url, appName, clientID, clientSecret, apiKey)
		if err != nil {
			return response.ErrorResponseStr(tokenError.Error(), fiber.StatusBadRequest, ctx)
		}
		/*-----------------------------------------------------------------
		 12. UPDATE TOKEN AND TOKEN TIME AND GET RESPONSE IF IS AVAILABLE
		-------------------------------------------------------------------*/
		dbResponse := repositories.EventHubConfigurationsRepository.UpdateTokenAndTokenTime(1, tokenResponse.Data.AccessToken, time.Now().Add(3*time.Hour))
		if dbResponse.RowsAffected == 0 {
			return response.ErrorResponse("Failed to update token time and token", fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 13. GET CONFIGURATIONS
		----------------------------------------------------------*/
		configurations = service.EventHubConfigurationsService.GetConfigurations()
		if configurations == nil {
			return response.ErrorResponseStr("No records found !", fiber.StatusBadRequest, ctx)
		}
	}
	/*---------------------------------------------------------
	 14. GET AZAMPAY CHECKOUT BASE URL
	----------------------------------------------------------*/
	checkoutBaseURL, checkoutBaseURLError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(6)
	if checkoutBaseURLError != nil {
		return response.ErrorResponseStr(checkoutBaseURLError.Error.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 15. PUSH USSD
	----------------------------------------------------------*/
	url := checkoutBaseURL + "/azampay/mno/checkout"
	pushUSSDResponse, pushUSSDError := helpers.AzamPayPushUSSD(
		url,
		request.PhoneNumber,
		strconv.FormatFloat(float64(request.Amount), 'f', -1, 32),
		request.Currency,
		request.OrderID,
		request.Provider,
		configurations.AzampayToken,
		apiKey,
	)
	if err != nil {
		return response.ErrorResponseStr(pushUSSDError.Error(), fiber.StatusBadRequest, ctx)
	}
	if !pushUSSDResponse.Success {
		return response.ErrorResponseStr(pushUSSDResponse.Message, fiber.StatusBadRequest, ctx)
	}
	request.TransactionID = pushUSSDResponse.TransactionID
	/*--------------------------------------------------------------------
	 16. ADD EVENT
	-----------------------------------------------------------------------*/
	err = service.EventHubPaymentService.AddPaymentTransaction(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse(pushUSSDResponse.Message, fiber.StatusOK, ctx)
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
	events, err := service.EventHubPaymentService.GetPaymentTransactions(pagination, request.Query, request.Status)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(events, fiber.StatusOK, ctx)
}

func (c eventHubPaymentController) UpdatePaymentStatus(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request payments.EventHubUpdatePaymentStatusRequest
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
	/*---------------------------------------------------------
	 04. UPDATE PAYMENT STATUS
	----------------------------------------------------------*/
	if request.Transactionstatus == constants.Failure {
		return response.ErrorResponse(request.Message, fiber.StatusBadRequest, ctx)
	}
	dbResponse := repositories.EventHubPaymentRepository.UpdatePaymentStatus(request.Reference, constants.Completed)
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to update payment status!", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse(request.Message, fiber.StatusOK, ctx)
}