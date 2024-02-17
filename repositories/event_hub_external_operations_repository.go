package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubExternalOperationsRepository = newEventHubExternalOperationsRepository()

type eventHubExternalOperationsRepository struct {
}

func newEventHubExternalOperationsRepository() eventHubExternalOperationsRepository {
	return eventHubExternalOperationsRepository{}
}

func (s eventHubExternalOperationsRepository) GetMicroServiceExternalOperationSetup(parameterID uint64) (string, *gorm.DB) {
	var externalOperationsSetup *models.EventHubExternalOperationsSetup
	userBD := db.Where("id = ?", parameterID).Find(&externalOperationsSetup)
	if userBD.RowsAffected == 0 {
		return "", userBD
	}
	return externalOperationsSetup.Value, nil
}
