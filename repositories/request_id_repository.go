package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubRequestIDRepository = newEventHubRequestIDRepository()

type eventHubRequestIDRepository struct{}

func newEventHubRequestIDRepository() eventHubRequestIDRepository {
	return eventHubRequestIDRepository{}
}

func (r eventHubRequestIDRepository) GetRequestID() (*models.EventHubRequestIDDTO, *gorm.DB) {
	var requestID *models.EventHubRequestIDDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryMicroServiceRequestIDActiveKey()).Find(&requestID)
	return requestID, urDB
}

func (r eventHubRequestIDRepository) GetMicroServiceExternalOperationsSetup(id uint64) (*models.EventHubExternalOperationsSetup, *gorm.DB) {
	var externalOperationsSetup *models.EventHubExternalOperationsSetup
	urDB := db.Find(&externalOperationsSetup, id)
	return externalOperationsSetup, urDB
}

func (r eventHubRequestIDRepository) UpdateSetupValue(id int, value string) *gorm.DB {
	return db.Model(&models.EventHubExternalOperationsSetup{}).Where("id = ?", id).Update("value", value)
}
