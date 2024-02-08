package seeders

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

func EventHubExternalOperationSetupTableSeeder(db *gorm.DB) error {
	parameter1 := models.EventHubExternalOperationsSetup{Parameter: "Firebase FCM One Signal Notification URL", Value: "https://onesignal.com/api/v1/notifications"}

	allParameters := []models.EventHubExternalOperationsSetup{
		parameter1,
	}

	err := db.Create(allParameters).Error
	if err != nil {
		return err
	}
	return nil
}
