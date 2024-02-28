package seeders

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

func EventHubExternalOperationSetupTableSeeder(db *gorm.DB) error {
	parameter1 := models.EventHubExternalOperationsSetup{Parameter: "Firebase FCM One Signal Notification URL", Value: "https://onesignal.com/api/v1/notifications"}
	parameter2 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Sender ID", Value: "ALECOtr"}
	parameter3 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Single Destination Message URL", Value: "https://mshastra.com/sendurl.aspx "}
	parameter4 := models.EventHubExternalOperationsSetup{Parameter: "Next SMS Authorization token", Value: "Bm@90893$ "}
	parameter5 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Authenticator Base Url", Value: "https://authenticator.azampay.co.tz"}
	parameter6 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Checkout Base Url", Value: "https://checkout.azampay.co.tz"}
	parameter7 := models.EventHubExternalOperationsSetup{Parameter: "Azampay App Name", Value: "EVENT HUB TZ"}
	parameter8 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client ID", Value: "2d648311-1eba-4b1b-8e4d-b080b33046e9"}
	parameter9 := models.EventHubExternalOperationsSetup{Parameter: "Azampay Client Secret", Value: "Wp4QhqcKNONluPNojE2VqsJUJS2LWnFbOmHN1rUsJsltWgIcvSoNOyX4AVZdKwi9te/ypTXmq2AX10FRCH52kDk8WdLHgyrjlkEcIutVYY25ZWJ2rvUWNGVOoJvvLw6Qx2OmZbL3ZyxRFcu2H3deQ2cUjlG4FwESu08bqvMTtJEpztP1MXDpqXq6PGuS/5ev/tk/X+V+0PQN13dtKvqF/5TJMVjM97chS7Usf9aKxAOL1yl3J0Fi9NO8+9hmxYdF9cyRlDKb5aBxy7Luz139xfQ1SqeuLAJEtM7ZegPkSmEHt3RSrRema/JJ+bkZ0TFCR4pTWxYytd37PImYMo7wV4LUn2KiKrqqu8HjTJ13HAUMLLOrrxcUzXKkOvxY5diTVywj2jJLc8wHTbv22rMIdkHZCjle+vS2tZOUzKGr7O0BA4JMGUEooSg5KNbrepaI1B4yGw399R4eCCtL8evwvj1cQI+xgtzNP6RrWzYYR3lgHSr3/PQwy3LyHTdL+jcpl0QKxod04eq0gTSOGoULKGWdPorLWEZgoZUUoCaB6QpnJiEqTn78tGBNsnvFpHnxqFTOSoNOg0DlLw37bkimlcnlMJQO4H+4u/hhQviCnppeff1Q/hp6YpBkdtRsMI9B4u3nww/Ix2RjvqZU+lqo7ZS4j9lEZ/b25neHoVtAAZg="}
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
