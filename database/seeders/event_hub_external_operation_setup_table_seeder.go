package seeders

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

func EventHubExternalOperationSetupTableSeeder(db *gorm.DB) error {
	parameter1 := models.EventHubExternalOperationsSetup{Parameter: "Firebase FCM One Signal Notification URL", Value: "https://onesignal.com/api/v1/notifications"}
	parameter2 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Sender ID", Value: "EVENT HUB"}
	parameter3 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Single Destination Message URL", Value: "https://messaging-service.co.tz/api/sms/v1/text/single"}
	parameter4 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Authorization token", Value: "Ym9uaXBoYWNlbWtpbmRpQGdtYWlsLmNvbTpCbUA5MDg5Mw=="}

	allParameters := []models.EventHubExternalOperationsSetup{
		parameter1, parameter2, parameter3, parameter4,
	}

	err := db.Create(allParameters).Error
	if err != nil {
		return err
	}
	return nil
}
