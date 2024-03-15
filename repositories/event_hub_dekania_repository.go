package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubDekaniaRepository = newEventHubDekaniaRepository()

type eventHubDekaniaRepository struct {
}

func newEventHubDekaniaRepository() eventHubDekaniaRepository {
	return eventHubDekaniaRepository{}
}

func (r eventHubDekaniaRepository) AddDekania(dekania models.EventHubDekania) (models.EventHubDekania, *gorm.DB) {
	chR := db.Create(&dekania)
	return dekania, chR
}

func (r eventHubDekaniaRepository) GetAllDekania() ([]models.EventHubDekaniaDTO, *gorm.DB) {

	var dekania []models.EventHubDekaniaDTO
	cmDB := db.Raw("select id,dekania_name,created_at,updated_at " +
		"FROM event_hub_dekania").Find(&dekania)
	return dekania, cmDB
}

func (r eventHubDekaniaRepository) GetAllDekaniaByPagination(pagination models.Pagination, query string) (models.Pagination, *gorm.DB) {

	dekania, urDB := helpers.EventHubQueryBuilder.QueryAllDekania(pagination, query)

	return dekania, urDB
}

func (r eventHubDekaniaRepository) GetRegionWithId(id uint64) (*models.EventHubDekania, *gorm.DB) {
	var dekania *models.EventHubDekania
	sRDB := db.Find(&dekania, id)
	return dekania, sRDB
}

func (r eventHubDekaniaRepository) UpdateRegionWithId(dekania *models.EventHubDekania) *gorm.DB {
	sRDB := db.Save(&dekania)
	return sRDB
}

func (r eventHubDekaniaRepository) DeleteDekania(regionId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", regionId).Delete(models.EventHubDekania{})
	return sRDB
}
