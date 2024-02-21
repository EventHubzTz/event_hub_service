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
	parameter5 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Authenticator Base Url", Value: "https://authenticator-sandbox.azampay.co.tz"}
	parameter6 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Checkout Base Url", Value: "https://sandbox.azampay.co.tz"}
	parameter7 := models.EventHubExternalOperationsSetup{Parameter: "Azampay App Name", Value: ""}
	parameter8 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client ID", Value: ""}
	parameter9 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client Secret", Value: ""}

	allParameters := []models.EventHubExternalOperationsSetup{
		parameter1, parameter2, parameter3, parameter4, parameter5,
		parameter6, parameter7, parameter8, parameter9,
	}

	err := db.Create(allParameters).Error
	if err != nil {
		return err
	}
	return nil
}
