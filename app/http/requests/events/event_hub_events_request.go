package events

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubEventRequest struct {
	UserID             uint64    `json:"user_id" validate:"required"`
	EventName          string    `json:"event_name" validate:"required"`
	EventLocation      string    `json:"event_location" validate:"required"`
	EventTime          time.Time `json:"event_time" validate:"required"`
	EventDescription   string    `json:"event_description" validate:"required"`
	EventCapacity      int       `json:"event_capacity" validate:"required"`
	EventEntrance      float32   `json:"event_entrance" validate:"required"`
	EventCategoryID    uint64    `json:"event_category_id" validate:"required"`
	EventSubCategoryID uint64    `json:"event_sub_category_id" validate:"required"`
}

func (request EventHubEventRequest) ToModel() models.EventHubEvent {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubEvent{
		UserID:             request.UserID,
		EventName:          request.EventName,
		EventLocation:      request.EventLocation,
		EventTime:          request.EventTime,
		EventDescription:   request.EventDescription,
		EventCapacity:      request.EventCapacity,
		EventEntrance:      request.EventEntrance,
		EventCategoryID:    request.EventCategoryID,
		EventSubCategoryID: request.EventSubCategoryID,
	}
}

type EventHubUpdateEventRequest struct {
	models.IDRequest
	EventName          string    `json:"event_name" validate:"required"`
	EventLocation      string    `json:"event_location" validate:"required"`
	EventTime          time.Time `json:"event_time" validate:"required"`
	EventDescription   string    `json:"event_description" validate:"required"`
	EventCapacity      int       `json:"event_capacity" validate:"required"`
	EventEntrance      float32   `json:"event_entrance" validate:"required"`
	EventCategoryID    uint64    `json:"event_category_id" validate:"required"`
	EventSubCategoryID uint64    `json:"event_sub_category_id" validate:"required"`
}

func (request EventHubUpdateEventRequest) ToModel() models.EventHubEvent {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubEvent{
		EventName:          request.EventName,
		EventLocation:      request.EventLocation,
		EventTime:          request.EventTime,
		EventDescription:   request.EventDescription,
		EventCapacity:      request.EventCapacity,
		EventEntrance:      request.EventEntrance,
		EventCategoryID:    request.EventCategoryID,
		EventSubCategoryID: request.EventSubCategoryID,
	}
}

type EventHubEventImageRequest struct {
	EventID       uint64 `json:"event_id" validate:"required"`
	Image         string `json:"image"`
	AspectRatios  string `json:"aspect_ratios"`
	FileType      string `json:"file_type"`
	ImagePath     string `json:"image_path"`
	ImageStorage  string `json:"image_storage"`
	ThumbunailUrl string `json:"thumbunail_url"`
	VideoUrl      string `json:"video_url"`
}

func (request EventHubEventImageRequest) ToModel() models.EventHubEventImages {
	return models.EventHubEventImages{
		EventID:       request.EventID,
		ImageUrl:      request.ImagePath,
		AspectRatios:  request.AspectRatios,
		FileType:      request.FileType,
		ImageStorage:  request.ImageStorage,
		ThumbunailUrl: request.ThumbunailUrl,
		VideoUrl:      request.VideoUrl,
	}
}

type EventHubEventGetRequest struct {
	EventID uint64 `json:"id" validate:"required"`
}

type EventHubEventsGetsRequest struct {
	Query                string `json:"query"`
	Status               string `json:"status"`
	ProductCategoryID    uint64 `json:"event_category_id"`
	ProductSubCategoryID uint64 `json:"event_sub_category_id"`
	Limit                int    `json:"limit,omitempty"`
	Page                 int    `json:"page,omitempty"`
	Sort                 string `json:"sort,omitempty"`
}
