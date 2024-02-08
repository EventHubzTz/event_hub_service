package users

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubUserChangePasswordRequest struct {
	UserID          int64  `json:"user_id" validate:"required"`
	OldPassword     string `json:"old_password" validate:"required,min=6"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_new_password" validate:"required,min=6"`
}

func (request EventHubUserChangePasswordRequest) ToModel(hashedMessage string) models.EventHubUser {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO USER MODEL
	----------------------------------------------------------*/
	return models.EventHubUser{
		Password: hashedMessage,
	}
}
