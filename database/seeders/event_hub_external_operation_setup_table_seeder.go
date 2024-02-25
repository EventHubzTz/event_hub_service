package seeders

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

func EventHubExternalOperationSetupTableSeeder(db *gorm.DB) error {
	parameter1 := models.EventHubExternalOperationsSetup{Parameter: "Firebase FCM One Signal Notification URL", Value: "https://onesignal.com/api/v1/notifications"}
	parameter2 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Sender ID", Value: "EVENT HUB"}
	parameter3 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Single Destination Message URL", Value: "https://messaging-service.co.tz/api/sms/v1/text/single"}
	parameter4 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Authorization token", Value: "Yl9ta2luZGk6Qm1AOTA4OTM="}
	parameter5 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Authenticator Base Url", Value: "https://authenticator-sandbox.azampay.co.tz"}
	parameter6 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Checkout Base Url", Value: "https://sandbox.azampay.co.tz"}
	parameter7 := models.EventHubExternalOperationsSetup{Parameter: "Azampay App Name", Value: "EVENT HUB"}
	parameter8 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client ID", Value: "a0432bca-6253-4592-a687-ce5c5e02408d"}
	parameter9 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client Secret", Value: "NSZogwbovVT45OvAx1D3pbkCX28sEltJPHbl5iO5bvUpB9FwFeSR9xRk5SH5eMjAArAN2nxnrYqRVuEhjrLrWnfrP5jAC605h+mw1P3X+0WtKdLGZYezgJ80VctKodEndt70a52aGxTZs9lkJlfHliYCizTRjcqT0Awjvi7Zn7zESpgNXUIOgL+dP9ei6YVMHKD3GHkbv2rnkiQVbaDBG25AcQzh/amCIhmniCRJ2o4ADVvPEpizy3mZDLMenPg57GIQej81lbds9anVa01WX7vXTnk2WULt4dUykFFctHEo8oPydm6HlBle6qWd1en4ORKiBr33QfFZUow7MWzVqk1CxSymcCmGRR5wXXEHzahALkHmfLKixzhacCCD6JJp7VdEVjbKo4ujdYWY0Cl6dxARcPs/UYxqYmRslYPWzLSucW+xSKQzVBvPABFQOLWK4xss2F85TetvOpTjqKLUzn0LqIHTkeRn2mnu4+1KLkq4Fxhq9WfSclvIlIbvqHWEbzgbXvTfI8hhAl60450x6+S1XUJrYjTrUTmSPXo3KtqnPNLiCzqvhclI7lxLcevGuZBTL988h3kslZUeMe/nZqI7nqS8iQmQt30sEfUvtDPLOhZR2gUUVyy/b+eLx4Oi/+EuB93TsgfjjLKHXwg06ivA7g0Qa/EcwWubRqfF2o4="}
	parameter10 := models.EventHubExternalOperationsSetup{Parameter: "Azampay API-Key Token", Value: "da9f7f5f-8251-4007-93a2-29b3d8992a20"}

	allParameters := []models.EventHubExternalOperationsSetup{
		parameter1, parameter2, parameter3, parameter4, parameter5,
		parameter6, parameter7, parameter8, parameter9, parameter10,
	}

	err := db.Create(allParameters).Error
	if err != nil {
		return err
	}
	return nil
}
