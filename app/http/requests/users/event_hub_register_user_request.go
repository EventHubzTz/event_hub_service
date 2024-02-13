package users

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubRegisterUserRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email,unique=event_hub_users.email,min=3,max=50"`
	PhoneNumber string `json:"phone_no" validate:"required,unique=event_hub_users.phone_number,min=3,max=20,country_code=TZ"`
	Gender      string `json:"gender" validate:"required,max=10"`
	Password    string `json:"password" validate:"required,max=50"`
	Role        string `json:"role" validate:"required,max=30"`
	AppID       string `json:"app_id" validate:"max=12"`
}

func (request EventHubRegisterUserRequest) ToModel() models.EventHubUser {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO USER MODEL
	----------------------------------------------------------*/
	return models.EventHubUser{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Gender:      request.Gender,
		Password:    request.Password,
		Role:        request.Role,
	}
}

type EventHubUsersGetsRequest struct {
	Role string `json:"role" validate:"required"`
}

type EventHubResendOTPRequest struct {
	UserID      int64  `json:"user_id" validate:"required"`
	PhoneNumber string `json:"phone_no" validate:"required,unique=event_hub_users.phone_number,min=3,max=20,country_code=TZ"`
	AppID       string `json:"app_id" validate:"max=12"`
}

type EventHubGenerateForgotPasswordOtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,country_code=TZ"`
	AppID       string `json:"app_id" validate:"max=12"`
}

type EventHubVerifyOTPResetPasswordRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20,country_code=TZ"`
	OTP         string `json:"otp" validate:"required,max=6"`
	UserID      uint64 `json:"user_id" validate:"required"`
}

type EventHubUpdatePasswordRequest struct {
	OTP      string `json:"otp_code" validate:"required,max=6"`
	UserID   uint64 `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (request EventHubUpdatePasswordRequest) ToModel(hashedMessage string) models.EventHubUser {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO USER MODEL
	----------------------------------------------------------*/
	return models.EventHubUser{
		Password: hashedMessage,
	}
}

type EventHubVerifyPhoneNumberUsingOTPRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20,country_code=TZ"`
	OTP         string `json:"otp" validate:"required,max=6"`
}

func (request EventHubVerifyPhoneNumberUsingOTPRequest) ToModel() models.EventHubUserOTPCode {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO USER MODEL
	----------------------------------------------------------*/
	return models.EventHubUserOTPCode{
		OTP:   request.OTP,
		Phone: request.PhoneNumber,
	}
}
