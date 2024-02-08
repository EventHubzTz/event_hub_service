package repositories

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
)

var EventHubUserTokenRepository = newEventHubUserTokenRepository()

type eventHubUserTokenRepository struct {
}

func newEventHubUserTokenRepository() eventHubUserTokenRepository {
	return eventHubUserTokenRepository{}
}

func (_ eventHubUserTokenRepository) Create(userToken *models.EventHubUserToken) error {
	return db.Create(&userToken).Error
}

func (_ eventHubUserTokenRepository) Update(userToken *models.EventHubUserToken) error {
	return db.Save(&userToken).Error
}

func (_ eventHubUserTokenRepository) UpdateToken(userToken *models.EventHubUserToken) error {
	return db.Table("afya_app_users_tokens").Where("user_id", userToken.UserID).Update("token", userToken.Token).Error
}

func (_ eventHubUserTokenRepository) Delete(userToken *models.EventHubUserToken) error {
	return db.Delete(&userToken).Error
}

func (_ eventHubUserTokenRepository) GetUserTokenByUserId(userID uint64) *models.EventHubUserToken {
	var userToken *models.EventHubUserToken = nil
	db.Where("user_id=?", userID).Find(&userToken)
	return userToken
}

func (_ eventHubUserTokenRepository) GetUserTokenByUserIdOnCreate(userID uint64) (*models.EventHubUserToken, error) {
	var userToken *models.EventHubUserToken
	dbErr := db.Where("user_id=?", userID).Find(&userToken)
	if dbErr.RowsAffected == 0 {
		return nil, errors.New("User has no token")
	}
	return userToken, nil
}
