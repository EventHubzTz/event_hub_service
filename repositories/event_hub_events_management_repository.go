package repositories

import (
	"os"

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

func (r eventHubEventsManagementRepository) AddEventImage(eventImage models.EventHubEventImages) (models.EventHubEventImages, *gorm.DB) {
	chR := db.Create(&eventImage)
	return eventImage, chR
}

func (r eventHubEventsManagementRepository) GetEvents(pagination models.Pagination, role, query string, userID, eventCategoryId, eventSubCategoryId uint64) (models.Pagination, *gorm.DB) {

	events, urDB := helpers.EventHubQueryBuilder.QueryGetEvents(pagination, role, query, userID, eventCategoryId, eventSubCategoryId)

	return events, urDB
}

func (r eventHubEventsManagementRepository) GetEvent(eventID uint64) (models.EventHubEventDTO, *gorm.DB) {
	baseUrl := os.Getenv("APP_URL")

	var event models.EventHubEventDTO
	clDB := db.Table("event_hub_events as t1").
		Joins("LEFT JOIN event_hub_event_categories t2 on t1.event_category_id = t2.id").
		Joins("LEFT JOIN event_hub_event_subcategories t3 on t1.event_sub_category_id = t3.id").
		Joins("LEFT JOIN event_hub_users t4 on t1.user_id = t4.id").
		Select(
			"t1.*",
			"t2.event_category_name",
			"t3.event_sub_category_name",
			"CONCAT(t4.first_name, ' ', t4.last_name) as event_owner",
			"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		).
		Preload("EventFiles", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_hub_event_images").
				Select(
					"event_hub_event_images.*",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.image_url) ELSE event_hub_event_images.image_url END image_url",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.video_url) ELSE event_hub_event_images.video_url END video_url",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.thumbunail_url) ELSE event_hub_event_images.thumbunail_url END thumbunail_url",
				)
		})
	if eventID != 0 {
		clDB = clDB.Where("t1.id = ?", eventID)
	}
	clDB = clDB.Find(&event)
	return event, clDB
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

func (r eventHubEventsManagementRepository) DeleteEventImage(eventImageId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", eventImageId).Delete(models.EventHubEventImages{})
	return sRDB
}

func (r eventHubEventsManagementRepository) DeleteEvent(regionId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", regionId).Delete(models.EventHubEvent{})
	return sRDB
}

func (r eventHubEventsManagementRepository) FindProductImagesByProductID(eventID uint64) ([]models.EventHubEventImagesDTO, *gorm.DB) {
	var contentCoverImage []models.EventHubEventImagesDTO
	ccDB := db.Where("event_id = ?", eventID).Find(&contentCoverImage)
	return contentCoverImage, ccDB
}

func (r eventHubEventsManagementRepository) GetDashboardStatistics() (*models.EventHubDashboardStatisticsDTO, *gorm.DB) {
	var statistics *models.EventHubDashboardStatisticsDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryGetDashboardStatistics()).Find(&statistics)
	return statistics, urDB
}
