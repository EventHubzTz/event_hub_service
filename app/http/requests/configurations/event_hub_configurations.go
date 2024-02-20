package configurations

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubConfigurationsRequest struct {
	AzampayTokenGeneratedTime time.Time `json:"azampay_token_generated_time" validate:"required"`
	AzampayToken              string    `json:"azampay_token" validate:"required"`
}

func (request EventHubConfigurationsRequest) ToModel() models.EventHubConfigurations {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubConfigurations{
		AzampayTokenGeneratedTime: request.AzampayTokenGeneratedTime,
		AzampayToken:              request.AzampayToken,
	}
}
