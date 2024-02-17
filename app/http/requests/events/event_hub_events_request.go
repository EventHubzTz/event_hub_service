package events

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubEventRequest struct {
	EventName        string    `json:"event_name" validate:"required"`
	EventLocation    string    `json:"event_location" validate:"required"`
	EventTime        time.Time `json:"event_time" validate:"required"`
	EventDescription string    `json:"event_description" validate:"required"`
}

func (request EventHubEventRequest) ToModel() models.EventHubEvent {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubEvent{
		EventName:        request.EventName,
		EventLocation:    request.EventLocation,
		EventTime:        request.EventTime,
		EventDescription: request.EventDescription,
	}
}

type EventHubUpdateEventRequest struct {
	models.IDRequest
	EventName        string    `json:"event_name" validate:"required"`
	EventLocation    string    `json:"event_location" validate:"required"`
	EventTime        time.Time `json:"event_time" validate:"required"`
	EventDescription string    `json:"event_description" validate:"required"`
}

func (request EventHubUpdateEventRequest) ToModel() models.EventHubEvent {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubEvent{
		EventName:        request.EventName,
		EventLocation:    request.EventLocation,
		EventTime:        request.EventTime,
		EventDescription: request.EventDescription,
	}
}

type EventHubEventGetRequest struct {
	EventID uint64 `json:"id" validate:"required"`
}

type EventHubEventsGetsRequest struct {
	Role  string `json:"role"`
	Query string `json:"query"`
	Limit int    `json:"limit,omitempty"`
	Page  int    `json:"page,omitempty"`
	Sort  string `json:"sort,omitempty"`
}
