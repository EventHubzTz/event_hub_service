package repositories

import (
	"os"

	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubCategoriesSubCategoriesRepository = newEventCategoriesSubCategoriesRepository()

type eventHubCategoriesSubCategoriesRepository struct {
}

func newEventCategoriesSubCategoriesRepository() eventHubCategoriesSubCategoriesRepository {
	return eventHubCategoriesSubCategoriesRepository{}
}

func (r eventHubCategoriesSubCategoriesRepository) CreateEventCategory(eventCategory models.EventHubEventCategories) (models.EventHubEventCategories, *gorm.DB) {
	chR := db.Create(&eventCategory)
	return eventCategory, chR
}

func (r eventHubCategoriesSubCategoriesRepository) GetAllEventCategories() ([]models.EventHubEventCategoriesDTO, *gorm.DB) {
	baseUrl := os.Getenv("APP_URL")

	var eventCategories []models.EventHubEventCategoriesDTO
	cmDB := db.Raw("select id,event_category_name,CONCAT(CASE image_storage WHEN 'LOCAL' THEN '" + baseUrl + "' ELSE '' END, icon_url) as icon_url," +
		"event_category_color FROM afya_app_products_categories").Find(&eventCategories)
	return eventCategories, cmDB
}

func (r eventHubCategoriesSubCategoriesRepository) GetAllEventCategoriesByPagination(pagination models.Pagination, query string) (models.Pagination, *gorm.DB) {

	eventCategories, urDB := helpers.EventHubQueryBuilder.QueryAllEventCategories(pagination, query)

	return eventCategories, urDB
}

func (r eventHubCategoriesSubCategoriesRepository) GetEventCategoryWithId(id uint64) (*models.EventHubEventCategories, *gorm.DB) {
	var eventCategory *models.EventHubEventCategories
	sRDB := db.Find(&eventCategory, id)
	return eventCategory, sRDB
}

func (r eventHubCategoriesSubCategoriesRepository) UpdateEventCategoryWithId(eventCategory *models.EventHubEventCategories) *gorm.DB {
	sRDB := db.Save(&eventCategory)
	return sRDB
}

func (r eventHubCategoriesSubCategoriesRepository) DeleteEventCategory(eventCategoryId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", eventCategoryId).Delete(models.EventHubEventCategories{})
	return sRDB
}

func (r eventHubCategoriesSubCategoriesRepository) CreateEventSubCategory(eventSubCategory models.EventHubEventSubCategories) (models.EventHubEventSubCategories, *gorm.DB) {
	chR := db.Create(&eventSubCategory)
	return eventSubCategory, chR
}

func (r eventHubCategoriesSubCategoriesRepository) GetAllEventSubCategories(eventCategoryId uint64) ([]models.EventHubEventSubCategoriesDTO, *gorm.DB) {
	baseUrl := os.Getenv("APP_URL")

	var productsSubCategories []models.EventHubEventSubCategoriesDTO
	cmDB := db.Table("afya_app_event_subcategories as sub").
		Select("sub.*", "CONCAT(CASE image_storage WHEN 'LOCAL' THEN '"+baseUrl+"' ELSE '' END, icon_url) as icon_url").
		Where("event_category_id = ?", eventCategoryId).
		Find(&productsSubCategories)
	return productsSubCategories, cmDB
}

func (r eventHubCategoriesSubCategoriesRepository) GetAllEventSubCategoriesByPagination(pagination models.Pagination, eventCategoryId uint64, query string) (models.Pagination, *gorm.DB) {

	eventSubCategories, urDB := helpers.EventHubQueryBuilder.QueryAllEventSubCategories(pagination, eventCategoryId, query)

	return eventSubCategories, urDB
}

func (r eventHubCategoriesSubCategoriesRepository) GetEventSubCategoryWithId(id uint64) (*models.EventHubEventSubCategories, *gorm.DB) {
	var eventSubCategory *models.EventHubEventSubCategories
	sRDB := db.Find(&eventSubCategory, id)
	return eventSubCategory, sRDB
}

func (r eventHubCategoriesSubCategoriesRepository) UpdateEventSubCategoryWithId(eventCategory *models.EventHubEventSubCategories) *gorm.DB {
	sRDB := db.Save(&eventCategory)
	return sRDB
}

func (r eventHubCategoriesSubCategoriesRepository) DeleteEventSubCategory(eventSubCategoryId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", eventSubCategoryId).Delete(models.EventHubEventSubCategories{})
	return sRDB
}
