package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubConfigurationsService = newEventHubConfigurationsService()

type eventHubConfigurationsService struct {
}

func newEventHubConfigurationsService() eventHubConfigurationsService {
	return eventHubConfigurationsService{}
}

func (s eventHubConfigurationsService) AddConfiguration(configuration models.EventHubConfigurations) error {
	/*---------------------------------------------------------
	 01. ADD CONFIGURATION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	_, dbResponse := repositories.EventHubConfigurationsRepository.AddConfiguration(&configuration)
	if dbResponse.RowsAffected == 0 {
		return errors.New("failed to add configuration! ")
	}
	return nil
}

func (s eventHubConfigurationsService) GetConfigurations() *models.EventHubConfigurationsDTO {
	configurations, usDB := repositories.EventHubConfigurationsRepository.GetConfigurations()
	if usDB.RowsAffected == 0 {
		return nil
	}
	return configurations
}
