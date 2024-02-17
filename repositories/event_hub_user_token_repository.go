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

func (r eventHubUserTokenRepository) Create(userToken *models.EventHubUserToken) error {
	return db.Create(&userToken).Error
}

func (r eventHubUserTokenRepository) Update(userToken *models.EventHubUserToken) error {
	return db.Save(&userToken).Error
}

func (r eventHubUserTokenRepository) UpdateToken(userToken *models.EventHubUserToken) error {
	return db.Table("afya_app_users_tokens").Where("user_id", userToken.UserID).Update("token", userToken.Token).Error
}

func (r eventHubUserTokenRepository) Delete(userToken *models.EventHubUserToken) error {
	return db.Delete(&userToken).Error
}

func (r eventHubUserTokenRepository) GetUserTokenByUserId(userID uint64) *models.EventHubUserToken {
	var userToken *models.EventHubUserToken = nil
	db.Where("user_id=?", userID).Find(&userToken)
	return userToken
}

func (r eventHubUserTokenRepository) GetUserTokenByUserIdOnCreate(userID uint64) (*models.EventHubUserToken, error) {
	var userToken *models.EventHubUserToken
	dbErr := db.Where("user_id=?", userID).Find(&userToken)
	if dbErr.RowsAffected == 0 {
		return nil, errors.New("user has no token")
	}
	return userToken, nil
}
