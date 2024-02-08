package users

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubRegisterUserRequest struct {
	FirstName         string `json:"first_name" validate:"required"`
	LastName          string `json:"last_name" validate:"required"`
	PhoneNumber       string `json:"phone_no" validate:"required,unique=event_hub_users.phone_number,min=3,max=20,country_code=TZ"`
	Gender            string `json:"gender" validate:"required,max=10"`
	Password          string `json:"password" validate:"required,max=50"`
	Role              string `json:"role" validate:"required,max=30"`
	AgentBusStandName string `json:"agent_bus_stand_name" validate:"required"`
}

func (request EventHubRegisterUserRequest) ToModel() models.EventHubUser {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO USER MODEL
	----------------------------------------------------------*/
	return models.EventHubUser{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: request.PhoneNumber,
		Gender:      request.Gender,
		Password:    request.Password,
		Role:        request.Role,
	}
}

type KataTiketiUsersGetsRequest struct {
	Role string `json:"role" validate:"required"`
}
