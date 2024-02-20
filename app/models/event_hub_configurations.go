package models

import "time"

type EventHubConfigurations struct {
	ID
	AzampayTokenGeneratedTime time.Time `json:"azampay_token_generated_time" gorm:"not null"`
	AzampayToken              string    `json:"azampay_token" gorm:"not null"`
	Timestamp
}

type EventHubConfigurationsDTO struct {
	ID
	AzampayTokenGeneratedTime string `json:"azampay_token_generated_time"`
	AzampayToken              string `json:"azampay_token"`
	TimestampString
}

func (EventHubConfigurations) TableName() string {
	return tablePrefix + "configurations"
}
