package service

import (
	"crypto/rand"
	"errors"
	"io"

	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/ggwhite/go-masker"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var EventHubUsersManagementService = newEventHubUsersManagementService()

type eventHubUsersManagementService struct {
}

func newEventHubUsersManagementService() eventHubUsersManagementService {
	return eventHubUsersManagementService{}
}

func (s eventHubUsersManagementService) RegisterUser(user models.EventHubUser) (*models.EventHubUser, error) {
	/*---------------------------------------------------------
	 01. CREATE USER AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	createdUser, dbResponse := repositories.EventHubUsersManagementRepository.RegisterUser(&user)
	if dbResponse.RowsAffected == 0 {
		return nil, errors.New(dbResponse.Error.Error())
	}

	/*---------------------------------------------------------
	 02. ADD USER IN USER TOKEN TABLE
	----------------------------------------------------------*/
	return createdUser, EventHubUserTokenService.CreateUserTokenInDB(&user)
}

func (s eventHubUsersManagementService) GetUserByEmailPhone(emailPhone string) (*models.EventHubUser, error) {
	user, dbResponse := repositories.EventHubUsersManagementRepository.FindOneByEmailPhone(emailPhone)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return user, errors.New("User with " + emailPhone + " not found! ")
	}
	return user, nil
}

func (s eventHubUsersManagementService) GetSpecificUserDetailsUsingPhoneNumber(phoneNumber string) *models.EventHubUserDTO {
	user, usDB := repositories.EventHubUsersManagementRepository.FindUserUsingPhoneNumber(phoneNumber)
	if usDB.RowsAffected == 0 {
		return nil
	}
	return user
}

func (s eventHubUsersManagementService) GetSpecificUser(id uint64) *models.EventHubUser {
	user, usDB := repositories.EventHubUsersManagementRepository.FindOne(id)
	if usDB.RowsAffected == 0 {
		return nil
	}
	return user
}

func (s eventHubUsersManagementService) GetUsers(pagination models.Pagination, role, query string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	users, dbResponse := repositories.EventHubUsersManagementRepository.GetUsers(pagination, role, newQuery)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("users not found! ")
	}
	return users, nil
}

func (s eventHubUsersManagementService) ChangePassword(user models.EventHubUser, userID uint64) (string, error) {
	/*---------------------------------------------------------
	 01. INITIATING VALUES HOLDING MODEL VALUES BEFORE THE
	     UPDATES
	----------------------------------------------------------*/
	userFromDatabase, usDB := repositories.EventHubUsersManagementRepository.FindOne(userID)
	if usDB.RowsAffected == 0 {
		return "", errors.New("internal server error")
	}
	userFromDatabase.Password = user.Password
	/*---------------------------------------------------------
	 02. UPDATING USER DETAILS TO THE DATABASE
	----------------------------------------------------------*/
	_, usrDB := repositories.EventHubUsersManagementRepository.UpdateWithID(userID, userFromDatabase)
	if usrDB.RowsAffected == 0 {
		return "", errors.New("internal server error")
	}

	return "Password updated successful", nil
}

func (s eventHubUsersManagementService) SendOtpToUser(userID uint64, appID, phone string) {
	successCounter := 0
	errorCounter := 0
	/*---------------------------------------------------------
	 01. SAVE THE OTP CODE FOR VERIFICATION OF THE MOBILE
	     APPLICATION
	----------------------------------------------------------*/
	_, dbErr, body := SaveOTPCode(phone, appID, int(userID))
	if dbErr.RowsAffected == 0 {
		return
	}

	senderID, errSenderID := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(2)
	if errSenderID == nil {
		messageUrl, errMessageUrl := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(3)
		if errMessageUrl == nil {
			authorizationToken, errAuthorizationToken := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(4)
			if errAuthorizationToken == nil {

				response, _ := helpers.MobiSMSApi(senderID, messageUrl, authorizationToken, phone, body)
				// response, _ := helpers.EventHubClientRESTAPIHelper.SendOTPMessageToMobileUser(senderID, messageUrl, authorizationToken, phone, body)

				/*--------------------------------------------
					04. SAVING RETURNED RESPONSE INTO A TABLE
				----------------------------------------------*/
				otpCodeMessageResponse := models.EventHubOTPMessageResponse{Value: string(response)}
				_, usrDB := repositories.EventHubUsersManagementRepository.SaveUserOTPCodeMessageResponse(&otpCodeMessageResponse)
				if usrDB.RowsAffected == 0 {
					errorCounter++
				} else {
					successCounter++
				}
			} else {
				errorCounter++
			}
		} else {
			errorCounter++
		}
	} else {
		errorCounter++
	}
}

func (s eventHubUsersManagementService) VerifyMobileNumberOTPCOde(otpCode models.EventHubUserOTPCode, user *models.EventHubUser, ctx *fiber.Ctx) error {
	/*---------------------------------------------------------
	 01. CHECK IF THE OTP CODE DETAILS ARE CORRECT IN THE
	     DATABASE
	----------------------------------------------------------*/
	phoneAndCodeStatus := repositories.EventHubUsersManagementRepository.VerifyPhoneNumberAndOTPCode(otpCode)
	if !phoneAndCodeStatus {
		return errors.New("invalid Phone number OTP Code")
	}
	/*---------------------------------------------------------
	 02. CHECK IF THE PHONE NUMBER IS THE CORRECT ONE FOR THE
	     PARTICULAR USER IN THE DATABASE
	----------------------------------------------------------*/
	phoneStatus := repositories.EventHubUsersManagementRepository.VerifyUserPhoneNumber(otpCode.Phone, user.Id)
	if !phoneStatus {
		return errors.New("invalid Phone number")
	}
	/*---------------------------------------------------------
	 03. CHECK IF THE PHONE NUMBER IS INVALID PARTICULAR USER
	     IN THE DATABASE
	----------------------------------------------------------*/
	phoneInvalidStatus := repositories.EventHubUsersManagementRepository.VerifyUserPhoneNumberInvalidStatus(otpCode.Phone, user.Id)
	if !phoneInvalidStatus {
		type verificationResponse struct {
			Message     string `json:"message"`
			Role        string `json:"role"`
			Username    string `json:"username"`
			Error       bool   `json:"error"`
			PhoneNumber string `json:"phone_number"`
		}

		verificationRes := verificationResponse{
			Message:     "Phone number " + masker.String(masker.MID, otpCode.Phone) + " already verified!",
			PhoneNumber: otpCode.Phone,
			Error:       false,
			Role:        "normal user",
		}

		return ctx.Status(fiber.StatusOK).JSON(verificationRes)
		// resp :=
	}
	/*---------------------------------------------------------
	 03. UPDATE THE STATUS OF THE 'is_valid_phone_number' FOR
	     THE PARTICULAR USER
	----------------------------------------------------------*/
	err := repositories.EventHubUsersManagementRepository.UpdateUserPhoneNumberValidStatus(user)
	if err != nil {
		return err
	}
	return nil
}

func (s eventHubUsersManagementService) GenerateForgotPasswordOtp(phoneNumber, appSignature string, user *models.EventHubUserDTO) interface{} {
	userForgetPasswordOTPDetails := repositories.EventHubUsersManagementRepository.FindForgetPasswordOTPDetails(user.Id)
	/*----------------------------------------------------------
	 01 FETCHING THE MOBILE APPLICATION ID FROM THE DATABASE
	------------------------------------------------------------*/
	type ResponseStatus struct {
		Error   bool   `json:"ERROR"`
		Message string `json:"MESSAGE"`
	}
	otpCode := generateOTPCode()
	message := "Hello " + user.FirstName + " " + user.LastName + ", OTP Code " + otpCode + ".\nUse it within 5 minutes.\n" + appSignature
	if userForgetPasswordOTPDetails == nil {
		otpCodeForgotPassword := models.EventHubForgotPasswordOTP{
			UserID:    user.Id,
			OTP:       otpCode,
			Phone:     phoneNumber,
			Message:   message,
			IsOTPSent: "YES",
		}
		_, usrDB := repositories.EventHubUsersManagementRepository.CreateOTPCodeForForgotPassword(&otpCodeForgotPassword)
		if usrDB.RowsAffected == 0 {
			return ResponseStatus{
				Error:   true,
				Message: usrDB.Error.Error(),
			}
		}
		go func() {
			s.SendOtpToUser(user.Id, appSignature, phoneNumber)
		}()
		return ResponseStatus{
			Error:   false,
			Message: "OTP code is sent",
		}
	} else {
		userForgetPasswordOTPDetails.OTP = otpCode
		userForgetPasswordOTPDetails.Message = message
		userForgetPasswordOTPDetails.IsOTPSent = "YES"
		_, usDB := repositories.EventHubUsersManagementRepository.UpdateOTPCodeForForgotPassword(userForgetPasswordOTPDetails)
		if usDB.RowsAffected == 0 {
			return ResponseStatus{
				Error:   true,
				Message: "failed to update OTP Code for resetting password!",
			}
		}
		go func() {
			s.SendOtpToUser(user.Id, appSignature, phoneNumber)
		}()
		return ResponseStatus{
			Error:   false,
			Message: "OTP code is sent",
		}
	}
}

func (s eventHubUsersManagementService) VerifyOTPResetPassword(userID uint64, otp string, phoneNumber string) error {
	userForgetPasswordOTPDetails := repositories.EventHubUsersManagementRepository.FindForgetPasswordOTPDetails(userID)
	if userForgetPasswordOTPDetails == nil {
		return errors.New("reset Password OTP Verification failed! ")
	} else {
		if userForgetPasswordOTPDetails.OTP == otp && userForgetPasswordOTPDetails.Phone == phoneNumber {
			return nil
		} else {
			return errors.New("reset Password OTP Verification failed! ")
		}
	}
}

func (s eventHubUsersManagementService) UpdateProfile(user models.EventHubUser, userID uint64) (string, error) {
	updateMessage := "Password updated sucessful"
	/*---------------------------------------------------------
	 01. UPDATING USER DETAILS TO THE DATABASE
	----------------------------------------------------------*/
	_, usrDB := repositories.EventHubUsersManagementRepository.UpdateWithID(userID, &user)
	if usrDB.RowsAffected == 0 {
		return "", errors.New("internal server error")
	}

	return updateMessage, nil
}

func SaveOTPCode(phoneNumber, appID string, userId int) (*models.EventHubUserOTPCode, *gorm.DB, string) {
	otp := generateOTPCode()

	otpCode := models.EventHubUserOTPCode{
		OTP:   otp,
		Phone: phoneNumber,
	}
	otpRes, urDB := repositories.EventHubUsersManagementRepository.SaveUserOTPCode(&otpCode)
	if urDB.RowsAffected == 0 {
		return otpRes, urDB, ""
	}

	sms, Dbres := createUserOTPCodeMessage(uint64(userId), appID, otp)
	if Dbres.RowsAffected == 0 {
		return nil, Dbres, ""
	}

	// commands.AFYAAPPSendOtpMessageToUsersWhoHaveNoOtpMessage.SendOtpToUser(int64(userId),sm,request.PhoneNumber)

	return nil, Dbres, sms.Body
}

func generateOTPCode() string {
	var codeLength = 6
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, codeLength)
	n, err := io.ReadAtLeast(rand.Reader, b, codeLength)
	if n != codeLength {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func createUserOTPCodeMessage(userID uint64, appId string, otpCode string) (*models.EventHubOTPCodeMessage, *gorm.DB) {
	/*----------------------------------------------------------
	 01 FETCHING THE MOBILE APPLICATION ID FROM THE DATABASE
	------------------------------------------------------------*/
	// appSignature := repositories.EventHubExternalOperationsRepository.GetApplicationSignature()
	message := "Your Event Hub OTP Code is: " + otpCode + " " + appId
	otpCodeMessage := models.EventHubOTPCodeMessage{
		UserID: userID,
		Body:   message,
	}
	return repositories.EventHubUsersManagementRepository.SaveUserOTPCodeMessage(&otpCodeMessage)
}

func (c eventHubUsersManagementService) SendSms(phone, body string) error {
	senderID, errSenderID := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(2)
	if errSenderID == nil {
		messageUrl, errMessageUrl := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(3)
		if errMessageUrl == nil {
			authorizationToken, errAuthorizationToken := repositories.EventHubExternalOperationsRepository.GetMicroServiceExternalOperationSetup(4)
			if errAuthorizationToken == nil {
				_, err := helpers.MobiSMSApi(senderID, messageUrl, authorizationToken, phone, body)
				// _, err := helpers.EventHubClientRESTAPIHelper.SendOTPMessageToMobileUser(senderID, messageUrl, authorizationToken, phone, body)
				if err != nil {
					return err
				}
			} else {
				return errors.New("error get authorization token")
			}
		} else {
			return errors.New("error get message url")
		}
	} else {
		return errors.New("error get sender id")
	}
	return nil
}
