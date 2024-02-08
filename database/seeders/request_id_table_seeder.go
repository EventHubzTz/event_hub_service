package seeders

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

func EventHubRequestIDTableSeeder(db *gorm.DB) error {
	/*---------------------------------
	  01. CREATING REQUEST ID
	 ----------------------------------*/
	defaultRequestID := models.EventHubRequestID{
		RequestID: "XhoO2yoeISBAJja8AGuul0hYomoEkXKK",
	}

	err := db.Create(&defaultRequestID).Error
	if err != nil {
		return err
	}
	return nil
}
