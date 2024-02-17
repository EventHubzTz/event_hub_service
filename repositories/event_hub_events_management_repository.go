package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubEventsManagementRepository = newEventHubEventsManagementRepository()

type eventHubEventsManagementRepository struct {
}

func newEventHubEventsManagementRepository() eventHubEventsManagementRepository {
	return eventHubEventsManagementRepository{}
}

func (r eventHubEventsManagementRepository) AddEvent(event *models.EventHubEvent) (*models.EventHubEvent, *gorm.DB) {
	urDB := db.Create(&event)
	return event, urDB
}

func (r_ eventHubEventsManagementRepository) GetEvents(pagination models.Pagination, query string) (models.Pagination, *gorm.DB) {

	events, urDB := helpers.EventHubQueryBuilder.QueryGetEvents(pagination, query)

	return events, urDB
}

func (r eventHubEventsManagementRepository) GetEvent(id uint64) (*models.EventHubEventDTO, *gorm.DB) {
	var event *models.EventHubEventDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryEventDetails(), id).Find(&event)
	return event, urDB
}

func (r eventHubEventsManagementRepository) GetEventWithId(id uint64) (*models.EventHubEvent, *gorm.DB) {
	var event *models.EventHubEvent
	sRDB := db.Find(&event, id)
	return event, sRDB
}

func (r eventHubEventsManagementRepository) UpdateEventWithId(event *models.EventHubEvent) *gorm.DB {
	sRDB := db.Save(&event)
	return sRDB
}

func (r eventHubEventsManagementRepository) DeleteEvent(regionId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", regionId).Delete(models.EventHubEvent{})
	return sRDB
}
