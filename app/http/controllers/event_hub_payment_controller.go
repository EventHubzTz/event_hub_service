package controllers

import (
	"os"
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
	type PaymentTransactionData struct {
		Results       string `json:"results"`
		Message       string `json:"message"`
		TransactionID string `json:"transaction_id"`
	}

	var errorPaymentData PaymentTransactionData

	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF PUSH
	     USSD
	---------------------------------------------------------*/
	var request payments.EventHubPaymentRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if request.Amount > constants.MAXAMOUNT {
		return response.ErrorResponseStr("The amount is too large. Maximum is TZS 500,000/=", fiber.StatusBadRequest, ctx)
	}
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
	/*---------------------------------------------------------
	 04. GET CONFIGURATIONS
	----------------------------------------------------------*/
	configurations := service.EventHubConfigurationsService.GetConfigurations()
	if configurations == nil {
		errorPaymentData.Message = "No records found !"
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 05. GET AZAMPAY API KEY
	----------------------------------------------------------*/
	apiKey, apiKeyError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(10)
	if apiKeyError != nil {
		errorPaymentData.Message = apiKeyError.Error.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 06. CHECK IF TOKEN TIME HAS EXPIRED
	----------------------------------------------------------*/
	generatedTime, err := time.Parse(time.RFC3339, configurations.AzampayTokenGeneratedTime)
	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if time.Now().After(generatedTime) {
		/*---------------------------------------------------------
		 07. GET AZAMPAY AUTHENTICATOR BASE URL
		----------------------------------------------------------*/
		authenticatorBaseURL, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(5)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 08. GET AZAMPAY APP NAME
		----------------------------------------------------------*/
		appName, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(7)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 09. GET AZAMPAY CLIENT ID
		----------------------------------------------------------*/
		clientID, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(8)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 10. GET AZAMPAY CLIENT SECRET
		----------------------------------------------------------*/
		clientSecret, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(9)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 11. GET BEARER TOKEN
		----------------------------------------------------------*/
		url := authenticatorBaseURL + "/AppRegistration/GenerateToken"
		tokenResponse, tokenError := helpers.GenerateAzamPayToken(url, appName, clientID, clientSecret, apiKey)
		if tokenError != nil {
			errorPaymentData.Message = tokenError.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*-----------------------------------------------------------------
		 12. UPDATE TOKEN AND TOKEN TIME AND GET RESPONSE IF IS AVAILABLE
		-------------------------------------------------------------------*/
		dbResponse := repositories.EventHubConfigurationsRepository.UpdateTokenAndTokenTime(1, tokenResponse.Data.AccessToken, time.Now().Add(3*time.Hour))
		if dbResponse.RowsAffected == 0 {
			errorPaymentData.Message = "Failed to update token time and token"
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 13. GET CONFIGURATIONS
		----------------------------------------------------------*/
		configurations = service.EventHubConfigurationsService.GetConfigurations()
		if configurations == nil {
			errorPaymentData.Message = "No records found !"
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
	}
	/*---------------------------------------------------------
	 14. GET AZAMPAY CHECKOUT BASE URL
	----------------------------------------------------------*/
	checkoutBaseURL, checkoutBaseURLError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(6)
	if checkoutBaseURLError != nil {
		errorPaymentData.Message = checkoutBaseURLError.Error.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
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
	if pushUSSDError != nil {
		errorPaymentData.Message = pushUSSDError.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if !pushUSSDResponse.Success {
		errorPaymentData.Message = pushUSSDResponse.Message
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	request.TransactionID = pushUSSDResponse.TransactionID
	/*--------------------------------------------------------------------
	 16. ADD PAYMENT TRANSACTION
	-----------------------------------------------------------------------*/
	err = service.EventHubPaymentService.AddPaymentTransaction(request.ToModel())
	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}

	paymentData := PaymentTransactionData{
		Results:       pushUSSDResponse.Results,
		Message:       pushUSSDResponse.Message,
		TransactionID: request.TransactionID,
	}

	return response.DataListSuccessResponse(paymentData, fiber.StatusOK, ctx)
}

func (c eventHubPaymentController) VotingPushUSSD(ctx *fiber.Ctx) error {
	type PaymentTransactionData struct {
		Results       string `json:"results"`
		Message       string `json:"message"`
		TransactionID string `json:"transaction_id"`
	}

	var errorPaymentData PaymentTransactionData

	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF PUSH
	     USSD
	---------------------------------------------------------*/
	var request payments.EventHubVotingPaymentRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if request.TotalAmount > constants.MAXAMOUNT {
		return response.ErrorResponseStr("The amount is too large. Maximum is TZS 500,000/=", fiber.StatusBadRequest, ctx)
	}
	request.Currency = constants.Currency
	request.OrderID = utils.GenerateOrderId()
	request.Provider = utils.CheckMobileNetwork(request.PhoneNumber)
	request.TotalAmount = 1000
	if request.VoteNumbers > 0 {
		request.TotalAmount = request.TotalAmount * float32(request.VoteNumbers)
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
	 04. GET CONFIGURATIONS
	----------------------------------------------------------*/
	configurations := service.EventHubConfigurationsService.GetConfigurations()
	if configurations == nil {
		errorPaymentData.Message = "No records found !"
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 05. GET AZAMPAY API KEY
	----------------------------------------------------------*/
	apiKey, apiKeyError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(10)
	if apiKeyError != nil {
		errorPaymentData.Message = apiKeyError.Error.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 06. CHECK IF TOKEN TIME HAS EXPIRED
	----------------------------------------------------------*/
	generatedTime, err := time.Parse(time.RFC3339, configurations.AzampayTokenGeneratedTime)
	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if time.Now().After(generatedTime) {
		/*---------------------------------------------------------
		 07. GET AZAMPAY AUTHENTICATOR BASE URL
		----------------------------------------------------------*/
		authenticatorBaseURL, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(5)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 08. GET AZAMPAY APP NAME
		----------------------------------------------------------*/
		appName, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(7)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 09. GET AZAMPAY CLIENT ID
		----------------------------------------------------------*/
		clientID, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(8)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 10. GET AZAMPAY CLIENT SECRET
		----------------------------------------------------------*/
		clientSecret, err := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(9)
		if err != nil {
			errorPaymentData.Message = err.Error.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 11. GET BEARER TOKEN
		----------------------------------------------------------*/
		url := authenticatorBaseURL + "/AppRegistration/GenerateToken"
		tokenResponse, tokenError := helpers.GenerateAzamPayToken(url, appName, clientID, clientSecret, apiKey)
		if tokenError != nil {
			errorPaymentData.Message = tokenError.Error()
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*-----------------------------------------------------------------
		 12. UPDATE TOKEN AND TOKEN TIME AND GET RESPONSE IF IS AVAILABLE
		-------------------------------------------------------------------*/
		dbResponse := repositories.EventHubConfigurationsRepository.UpdateTokenAndTokenTime(1, tokenResponse.Data.AccessToken, time.Now().Add(3*time.Hour))
		if dbResponse.RowsAffected == 0 {
			errorPaymentData.Message = "Failed to update token time and token"
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
		/*---------------------------------------------------------
		 13. GET CONFIGURATIONS
		----------------------------------------------------------*/
		configurations = service.EventHubConfigurationsService.GetConfigurations()
		if configurations == nil {
			errorPaymentData.Message = "No records found !"
			return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
		}
	}
	/*---------------------------------------------------------
	 14. GET AZAMPAY CHECKOUT BASE URL
	----------------------------------------------------------*/
	checkoutBaseURL, checkoutBaseURLError := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(6)
	if checkoutBaseURLError != nil {
		errorPaymentData.Message = checkoutBaseURLError.Error.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 15. PUSH USSD
	----------------------------------------------------------*/
	url := checkoutBaseURL + "/azampay/mno/checkout"
	pushUSSDResponse, pushUSSDError := helpers.AzamPayPushUSSD(
		url,
		request.PhoneNumber,
		strconv.FormatFloat(float64(request.TotalAmount), 'f', -1, 32),
		request.Currency,
		request.OrderID,
		request.Provider,
		configurations.AzampayToken,
		apiKey,
	)
	if pushUSSDError != nil {
		errorPaymentData.Message = pushUSSDError.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	if !pushUSSDResponse.Success {
		errorPaymentData.Message = pushUSSDResponse.Message
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}
	request.TransactionID = pushUSSDResponse.TransactionID
	/*--------------------------------------------------------------------
	 16. ADD PAYMENT TRANSACTION
	-----------------------------------------------------------------------*/
	err = service.EventHubPaymentService.AddVotingPaymentTransaction(request.ToModel())
	if err != nil {
		errorPaymentData.Message = err.Error()
		return response.DataListErrorResponse(errorPaymentData, fiber.StatusBadRequest, ctx)
	}

	paymentData := PaymentTransactionData{
		Results:       pushUSSDResponse.Results,
		Message:       pushUSSDResponse.Message,
		TransactionID: request.TransactionID,
	}

	return response.DataListSuccessResponse(paymentData, fiber.StatusOK, ctx)
}

func (c eventHubPaymentController) GetVotingPaymentTransactions(ctx *fiber.Ctx) error {
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
	paymentTransactions, err := service.EventHubPaymentService.GetVotingPaymentTransactions(pagination, request.Query, request.Status)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(paymentTransactions, fiber.StatusOK, ctx)
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
	user := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	paymentTransactions, err := service.EventHubPaymentService.GetPaymentTransactions(pagination, user.Role, request.Query, request.Status, request.PhoneNumber, user.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(paymentTransactions, fiber.StatusOK, ctx)
}

func (c eventHubPaymentController) GetTransactionByTransactionID(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request payments.EventHubGetTransactionRequest
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
	 04. GET TRANSACTION
	----------------------------------------------------------*/
	transaction := repositories.EventHubPaymentRepository.GetTransactionByTransactionID(request.TransactionID)
	if transaction == nil {
		return response.ErrorResponse("Transaction does not exist in the system", fiber.StatusInternalServerError, ctx)
	}

	return response.InternalServiceDataResponse(transaction, fiber.StatusOK, ctx)
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
	 04. CHECK IF SUBSCRIPTION NOT COMPLETED
	----------------------------------------------------------*/
	if request.Transactionstatus == constants.Failure {
		return response.ErrorResponse(request.Message, fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 05. GET VOTING TRANSACTION
	----------------------------------------------------------*/
	votingTransaction := repositories.EventHubPaymentRepository.GetVotingTransactionByTransactionID(request.Reference)
	if votingTransaction != nil {
		/*---------------------------------------------------------
		 06. UPDATE PAYMENT STATUS
		----------------------------------------------------------*/
		dbResponse := repositories.EventHubPaymentRepository.UpdateVotingPaymentStatus(request.Reference, constants.Completed)
		if dbResponse.RowsAffected == 0 {
			return response.ErrorResponse("Failed to update payment status!", fiber.StatusBadRequest, ctx)
		}
		/*-----------------------------------------------------------------
		 07. SEND MESSAGE
		-------------------------------------------------------------------*/
		request.Message = "Malipo yako ya TZS " + strconv.Itoa(int(votingTransaction.TotalAmount)) + " kwa kupiga kura " + strconv.Itoa(int(votingTransaction.VoteNumbers)) + " katika kipengele cha " + votingTransaction.Category + " yamekamilika. Tarehe: " + date_utils.GetNowString() + ", Asante kwa kushiriki katika kutambua vipaji vya Tamthilia zetu.Kumbukumbu Namba: " + request.Reference
		err = service.EventHubUsersManagementService.SendSms(votingTransaction.PhoneNumber, request.Message)
		if err != nil {
			return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
		}
		/*-----------------------------------------------------------------
		 08. CALL VOTE URL
		-------------------------------------------------------------------*/
		url := os.Getenv("VOTE_URL")
		_, err = helpers.Vote(url, *votingTransaction)
		if err != nil {
			return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
		}

		return response.SuccessResponse(request.Message, fiber.StatusOK, ctx)
	}
	/*---------------------------------------------------------
	 09. GET TRANSACTION
	----------------------------------------------------------*/
	transaction := repositories.EventHubPaymentRepository.GetVotingTransactionByTransactionID(request.Reference)
	if transaction != nil {
		/*---------------------------------------------------------
		 10. UPDATE PAYMENT STATUS
		----------------------------------------------------------*/
		dbResponse := repositories.EventHubPaymentRepository.UpdatePaymentStatus(request.Reference, constants.Completed)
		if dbResponse.RowsAffected == 0 {
			return response.ErrorResponse("Failed to update payment status!", fiber.StatusBadRequest, ctx)
		}
	}
	/*-----------------------------------------------------------------
	 11. UPDATE TANZANIA GOSPEL MUSIC AWARDS PAYMENTS
	-------------------------------------------------------------------*/
	url := os.Getenv("UPDATE_TGMA_PAYMENT_URL")
	payment, err := helpers.UpdateTGMAPayments(url, request.Reference)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	} else if payment.Error {
		return response.ErrorResponse(payment.Message, fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse(request.Message, fiber.StatusOK, ctx)
}
