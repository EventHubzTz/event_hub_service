package controllers

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/http/requests/users"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubUsersManagementController = newEventHubUsersManagementController()

type eventHubUsersManagementController struct {
}

func newEventHubUsersManagementController() eventHubUsersManagementController {
	return eventHubUsersManagementController{}
}

func (_ eventHubUsersManagementController) RegisterUser(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF REGISTERING
	     USER
	---------------------------------------------------------*/
	var request users.EventHubRegisterUserRequest
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
	 04. HASH THE REQUEST PASSWORD AND RETURN BCRYPT HASHED PASSWORD
	----------------------------------------------------------*/
	request.Password = service.EventHubAuthenticationService.HashPassword(request.Password)
	/*---------------------------------------------------------
	 05. CREATE USER AND GET ERROR IF IS AVAILABLE
	----------------------------------------------------------*/
	_, error := service.EventHubUsersManagementService.RegisterUser(request.ToModel())
	/*---------------------------------------------------------
	 06. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if error != nil {
		return response.ErrorResponse(error.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 07. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.SuccessResponse("User registered successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (_ eventHubUsersManagementController) LoginUser(ctx *fiber.Ctx) error {

	type RegisterData struct {
		Id       int64  `json:"id"`
		Token    string `json:"token"`
		FullName string `json:"full_name"`
		Role     string `json:"role"`
	}

	type LoginRequest struct {
		Phone    string `json:"phone" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF CREATING
	     DOCTOR
	---------------------------------------------------------*/
	var request LoginRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)
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
	 04. AUTHENTICATE USER AND GET ERROR IF IS AVAILABLE AND SUCCESS MESSAGE
	-------------------------------------------------------------------*/
	user, err := service.EventHubUsersManagementService.GetUserByPhone(request.Phone)

	if err != nil || user.Active == false {
		return response.ErrorResponse("Bad credentials", fiber.StatusBadRequest, ctx)
	}
	/*-----------------------------------------------------------z
	 06. CHECK IF THE PASSWORD IS CORRECT
	------------------------------------------------------------*/
	isCurrentLoginPassword := service.EventHubAuthenticationService.CheckPasswordHash(request.Password, user.Password)
	if !isCurrentLoginPassword {
		return response.ErrorResponse("Bad credentials", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 05. SAVE THE TOKEN CODE FOR VERIFICATION OF THE SESSIONS
	----------------------------------------------------------*/
	_, token, _ := service.EventHubUserTokenService.GenerateToken(user)
	year := time.Hour * 24 * 365
	now := time.Now()
	futureDate := now.Add(year)
	tokenErr := service.EventHubUserTokenService.UpdateUserTokenInDB(user.Id, futureDate, token)
	if tokenErr != nil {
		return response.ErrorResponse("Failed to save token", fiber.StatusInternalServerError, ctx)
	}

	/*---------------------------------------------------------
	 06. DATA FOR SENDING RESPONCES
	----------------------------------------------------------*/
	regData := RegisterData{
		Id:       int64(user.Id),
		Token:    repositories.EventHubUserTokenRepository.GetUserTokenByUserId(user.Id).Token,
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
	}

	return response.DataListSuccessResponse(regData, fiber.StatusOK, ctx)
}

func (_ eventHubUsersManagementController) GetUsers(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF CREATING
	     CAR
	---------------------------------------------------------*/
	var request users.KataTiketiUsersGetsRequest
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
	/*-----------------------------------------------------------------
	 04. GET ALL USERS
	-------------------------------------------------------------------*/
	users, err := service.EventHubUsersManagementService.GetUsers(request.Role)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	return response.DataListSuccessResponse(users, fiber.StatusOK, ctx)
}

func (_ eventHubUsersManagementController) ChangePassword(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF UPDATING USER
	     PROFILE IMAGE
	---------------------------------------------------------*/
	var request users.EventHubUserChangePasswordRequest
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
	/*--------------------------------------------------------------------
	 04. CHECK IF USER EXIST
	-----------------------------------------------------------------------*/
	user := repositories.EventHubUsersManagementRepository.FindUserById(uint64(request.UserID))
	if user == nil {
		return response.ErrorResponse("User does not exist in the system", fiber.StatusInternalServerError, ctx)
	}
	/*----------------------------------------------------------
	 05. CHECK IF THE 'NEW PASSWORD' EQUALS TO THE 'CONFIRM NEW
	    PASSWORD'
	------------------------------------------------------------*/
	if request.NewPassword != request.ConfirmPassword {
		return response.ErrorResponse("Invalid Request. New Password do not match with the Confirm Password", fiber.StatusBadRequest, ctx)
	}
	/*-----------------------------------------------------------z
	 06. CHECK IF THE 'CURRENT PASSWORD' IS CORRECT
	------------------------------------------------------------*/
	isCurrentLoginPassword := service.EventHubAuthenticationService.CheckPasswordHash(request.OldPassword, user.Password)
	if !isCurrentLoginPassword {
		return response.ErrorResponse("Invalid request. Invalid Current Password", fiber.StatusBadRequest, ctx)
	}
	/*-----------------------------------------------------------
	 07. CHECK IF THE 'NEW PASSWORD' EQUALS TO THE 'OLD PASSWORD'
	------------------------------------------------------------*/
	if request.NewPassword == request.OldPassword {
		return response.ErrorResponse("Invalid request. You can't use your old password.", fiber.StatusBadRequest, ctx)
	}
	request.NewPassword = service.EventHubAuthenticationService.HashPassword(request.NewPassword)
	message, _err := service.EventHubUsersManagementService.ChangePassword(request.ToModel(request.NewPassword), user.Id)
	/*---------------------------------------------------------
	 08. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if _err != nil {
		return response.ErrorResponse(_err.Error(), fiber.StatusBadRequest, ctx)
	}
	message = "User password updated successful on " + date_utils.GetNowString() + "" + message
	/*---------------------------------------------------------
	 9. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.SuccessResponse(message, fiber.StatusOK, ctx)
}
