package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubUsersManagementRepository = newEventHubUsersManagementRepository()

type eventHubUsersManagementRepository struct {
}

func newEventHubUsersManagementRepository() eventHubUsersManagementRepository {
	return eventHubUsersManagementRepository{}
}

func (_ eventHubUsersManagementRepository) RegisterUser(user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Create(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) FindOneByPhone(phone string) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Model(models.EventHubUser{}).Where("phone_number = ?", phone).Find(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) GetUsers(role string) ([]models.KataTiketiUserDTO, *gorm.DB) {
	var renting []models.KataTiketiUserDTO
	sRDB := db.Raw(helpers.EventHubQueryBuilder.QueryGetUsers(role), role).Find(&renting)
	return renting, sRDB
}

func (_ eventHubUsersManagementRepository) FindUserById(userId uint64) *models.EventHubUser {
	var user *models.EventHubUser
	dbErr := db.Where("id = ?", userId).Find(&user)
	if dbErr.RowsAffected == 0 {
		return nil
	}
	return user
}

func (_ eventHubUsersManagementRepository) FindOne(id uint64) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryUserDetails(), id).Find(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) UpdateWithID(id uint64, user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Model(models.EventHubUser{}).Where("id = ?", id).Updates(&user)
	return user, urDB
}
