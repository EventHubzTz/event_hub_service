package repositories

import (
	"time"

	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubConfigurationsRepository = newEventHubConfigurationsRepository()

type eventHubConfigurationsRepository struct {
}

func newEventHubConfigurationsRepository() eventHubConfigurationsRepository {
	return eventHubConfigurationsRepository{}
}

func (r eventHubConfigurationsRepository) AddConfiguration(configuration *models.EventHubConfigurations) (*models.EventHubConfigurations, *gorm.DB) {
	urDB := db.Create(&configuration)
	return configuration, urDB
}

func (r eventHubConfigurationsRepository) GetConfigurations() (*models.EventHubConfigurationsDTO, *gorm.DB) {
	var configurations *models.EventHubConfigurationsDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryConfigurations()).Find(&configurations)
	return configurations, urDB
}

func (r eventHubConfigurationsRepository) UpdateTokenAndTokenTime(id uint64, token string, tokenTime time.Time) *gorm.DB {

	urDB := db.Model(models.EventHubConfigurations{}).Where("id = ? ", id).Updates(map[string]interface{}{
		"azampay_token":                token,
		"azampay_token_generated_time": tokenTime,
	})
	return urDB
}
