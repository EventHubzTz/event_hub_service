package controllers

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/http/requests/users"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/ggwhite/go-masker"
	"github.com/gofiber/fiber/v2"
)

var EventHubUsersManagementController = newEventHubUsersManagementController()

type eventHubUsersManagementController struct {
}

func newEventHubUsersManagementController() eventHubUsersManagementController {
	return eventHubUsersManagementController{}
}

type RegisterData struct {
	Id       int64  `json:"id"`
	Token    string `json:"token"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (c eventHubUsersManagementController) RegisterUser(ctx *fiber.Ctx) error {
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
	user, error := service.EventHubUsersManagementService.RegisterUser(request.ToModel())
	/*---------------------------------------------------------
	 06. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if error != nil {
		return response.ErrorResponse(error.Error(), fiber.StatusBadRequest, ctx)
	}
	go func() {
		service.EventHubUsersManagementService.SendOtpToUser(user.Id, request.AppID, request.PhoneNumber)
	}()
	/*---------------------------------------------------------
	 07. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	regData := RegisterData{
		Id:       int64(user.Id),
		Token:    repositories.EventHubUserTokenRepository.GetUserTokenByUserId(user.Id).Token,
		FullName: user.FirstName + " " + user.LastName,
		Role:     user.Role,
	}

	return response.DataListSuccessResponse(regData, fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) ResendOTPCode(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF REGISTERING
	     USER
	---------------------------------------------------------*/
	var request users.EventHubResendOTPRequest
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

	user := service.EventHubUserTokenService.GetUserFromLocal(ctx)

	go func() {
		service.EventHubUsersManagementService.SendOtpToUser(user.Id, request.AppID, request.PhoneNumber)
	}()

	return response.SuccessResponse("OTP code sent successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) VerifyPhoneNumberUsingOTP(ctx *fiber.Ctx) error {
	type verificationResponse struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF VERIFICATION
	     OF THE PHONE NUMBER USING OTP
	---------------------------------------------------------*/
	var request users.EventHubVerifyPhoneNumberUsingOTPRequest
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
	 04. VERIFICATION OF THE PHONE NUMBER IF ALL THE INPUT
	     VALIDATION PASSED
	----------------------------------------------------------*/
	user := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	err = service.EventHubUsersManagementService.VerifyMobileNumberOTPCOde(request.ToModel(), user, ctx)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 07. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	message := "Phone number " + masker.String(masker.MID, request.PhoneNumber) + " verified successfully on " + date_utils.GetNowString()
	/*---------------------------------------------------------
	 08. LOG USER ACTIVITY TO THE SYSTEM
	----------------------------------------------------------*/
	verificationRes := verificationResponse{
		Error:   false,
		Message: message,
	}

	return ctx.Status(fiber.StatusOK).JSON(verificationRes)
}

func (c eventHubUsersManagementController) LoginUser(ctx *fiber.Ctx) error {
	type LoginRequest struct {
		EmailPhone string `json:"email_phone" validate:"required"`
		Password   string `json:"password" validate:"required"`
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
	user, err := service.EventHubUsersManagementService.GetUserByEmailPhone(request.EmailPhone)

	if err != nil || !user.Active {
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

func (c eventHubUsersManagementController) GetUsers(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request users.EventHubUsersGetsRequest
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
	 04. GET USER AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	doctorData, err := service.EventHubUsersManagementService.GetUsers(pagination, request.Role, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(doctorData, fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) GetUser(ctx *fiber.Ctx) error {
	return response.InternalServiceDataResponse(service.EventHubUserTokenService.GetUserFromLocal(ctx), fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) ChangePassword(ctx *fiber.Ctx) error {
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
	userFromLocal := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	user := repositories.EventHubUsersManagementRepository.FindUserById(userFromLocal.Id)
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

func (c eventHubUsersManagementController) GenerateForgotPasswordOtp(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST FOR GENERATING
	     OTP CODE
	---------------------------------------------------------*/
	var request users.EventHubGenerateForgotPasswordOtpRequest
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
	 04. GET THE USER FROM THE DATABASE USING THE PHONE NUMBER
	----------------------------------------------------------*/
	userFromDB := service.EventHubUsersManagementService.GetSpecificUserDetailsUsingPhoneNumber(request.PhoneNumber)
	if userFromDB == nil {
		return response.ErrorResponse("Invalid status. User details not found in the system with a particular phone number", fiber.StatusBadRequest, ctx)
	}
	/*-------------------------------------------------------
	 05. GENERATE FORGOT PASSWORD OTP CODE
	---------------------------------------------------------*/
	return response.MapDataResponse(service.EventHubUsersManagementService.GenerateForgotPasswordOtp(request.PhoneNumber, request.AppID, userFromDB), fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) VerifyOTPResetPassword(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST VERIFY OTP
	     REST PASSWORD
	---------------------------------------------------------*/
	var request users.EventHubVerifyOTPResetPasswordRequest
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
	 04. GET THE USER FROM THE DATABASE USING THE PHONE NUMBER
	----------------------------------------------------------*/
	userFromLocal := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	userFromDB := service.EventHubUsersManagementService.GetSpecificUserDetailsUsingPhoneNumber(request.PhoneNumber)
	if userFromDB == nil {
		return response.ErrorResponse("Invalid status. User details not found in the system with a particular phone number", fiber.StatusBadRequest, ctx)
	}
	/*-----------------------------------------------
	 05. VERIFY OTP FOR RESETTING PASSWORD
	------------------------------------------------*/
	_err := service.EventHubUsersManagementService.VerifyOTPResetPassword(userFromLocal.Id, request.OTP, request.PhoneNumber)
	/*---------------------------------------------------------
	 06. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if _err != nil {
		return response.ErrorResponse(_err.Error(), fiber.StatusBadRequest, ctx)
	}
	message := "Password Verified successfully on " + date_utils.GetNowString()
	return response.SuccessResponse(message, fiber.StatusOK, ctx)
}

func (c eventHubUsersManagementController) UpdatePassword(ctx *fiber.Ctx) error {
	/*--------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST UPDATING PASSWORD
	---------------------------------------------------------*/
	var request users.EventHubUpdatePasswordRequest
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
	 04. GET THE USER FROM THE DATABASE USING USER ID
	----------------------------------------------------------*/
	userFromLocal := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	userFromDB := service.EventHubUsersManagementService.GetSpecificUser(userFromLocal.Id)
	if userFromDB == nil {
		return response.ErrorResponse("Invalid status. User details not found in the system", fiber.StatusBadRequest, ctx)
	}
	/*-----------------------------------------------
	 05. VERIFY OTP FOR RESETTING PASSWORD
	------------------------------------------------*/
	_err := service.EventHubUsersManagementService.VerifyOTPResetPassword(userFromLocal.Id, request.OTP, userFromDB.PhoneNumber)
	/*---------------------------------------------------------
	 06. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if _err != nil {
		return response.ErrorResponse(_err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 07. HASH PASSWORD AND UPDATE USER PASSWORD
	----------------------------------------------------------*/
	request.Password = service.EventHubAuthenticationService.HashPassword(request.Password)
	message, _err := service.EventHubUsersManagementService.UpdateProfile(request.ToModel(request.Password), userFromDB.Id)
	/*---------------------------------------------------------
	 08. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if _err != nil {
		return response.ErrorResponse(_err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 09. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.SuccessResponse(message, fiber.StatusOK, ctx)
}
