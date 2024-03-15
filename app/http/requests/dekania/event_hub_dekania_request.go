package dekania

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubDekaniaRequest struct {
	DekaniaName string `json:"dekania_name" validate:"required"`
}

type EventHubDekaniaUpdateRequest struct {
	models.IDRequest
	DekaniaName string `json:"dekania_name" validate:"required"`
}

type EventHubDekaniaGetRequest struct {
	DekaniaID uint64 `json:"id" validate:"required"`
}

func (request EventHubDekaniaRequest) ToModel() models.EventHubDekania {
	return models.EventHubDekania{
		DekaniaName: request.DekaniaName,
	}
}

func (request EventHubDekaniaUpdateRequest) ToModel() models.EventHubDekania {
	return models.EventHubDekania{
		DekaniaName: request.DekaniaName,
	}
}
