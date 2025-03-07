package dekania

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubRegionRequest struct {
	RegionName string `json:"region_name" validate:"required"`
}

type EventHubRegionUpdateRequest struct {
	models.IDRequest
	RegionName string `json:"region_name" validate:"required"`
}

type EventHubRegionGetRequest struct {
	DekaniaID uint64 `json:"id" validate:"required"`
}

func (request EventHubRegionRequest) ToModel() models.EventHubRegion {
	return models.EventHubRegion{
		RegionName: request.RegionName,
	}
}

func (request EventHubRegionUpdateRequest) ToModel() models.EventHubRegion {
	return models.EventHubRegion{
		RegionName: request.RegionName,
	}
}
