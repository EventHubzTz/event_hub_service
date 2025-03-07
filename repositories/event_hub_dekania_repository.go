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

func (r eventHubDekaniaRepository) AddRegion(region models.EventHubRegion) (models.EventHubRegion, *gorm.DB) {
	chR := db.Create(&region)
	return region, chR
}

func (r eventHubDekaniaRepository) GetAllRegions() ([]models.EventHubRegionDTO, *gorm.DB) {

	var region []models.EventHubRegionDTO
	cmDB := db.Raw("select id,dekania_name,created_at,updated_at " +
		"FROM event_hub_dekania").Find(&region)
	return region, cmDB
}

func (r eventHubDekaniaRepository) GetAllRegionsByPagination(pagination models.Pagination, query string) (models.Pagination, *gorm.DB) {

	region, urDB := helpers.EventHubQueryBuilder.QueryAllDekania(pagination, query)

	return region, urDB
}

func (r eventHubDekaniaRepository) GetRegionWithId(id uint64) (*models.EventHubRegion, *gorm.DB) {
	var region *models.EventHubRegion
	sRDB := db.Find(&region, id)
	return region, sRDB
}

func (r eventHubDekaniaRepository) UpdateRegionWithId(region *models.EventHubRegion) *gorm.DB {
	sRDB := db.Save(&region)
	return sRDB
}

func (r eventHubDekaniaRepository) DeleteRegion(regionId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", regionId).Delete(models.EventHubRegion{})
	return sRDB
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

func (r eventHubDekaniaRepository) GetDekaniaWithId(id uint64) (*models.EventHubDekania, *gorm.DB) {
	var dekania *models.EventHubDekania
	sRDB := db.Find(&dekania, id)
	return dekania, sRDB
}

func (r eventHubDekaniaRepository) UpdateDekaniaWithId(dekania *models.EventHubDekania) *gorm.DB {
	sRDB := db.Save(&dekania)
	return sRDB
}

func (r eventHubDekaniaRepository) DeleteDekania(regionId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", regionId).Delete(models.EventHubDekania{})
	return sRDB
}
