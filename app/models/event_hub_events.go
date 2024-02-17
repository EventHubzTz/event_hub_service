package models

import "time"

type EventHubEvent struct {
	ID
	EventName        string    `json:"event_name" gorm:"not null"`
	EventLocation    string    `json:"event_location" gorm:"not null"`
	EventTime        time.Time `json:"event_time" gorm:"not null"`
	EventDescription string    `json:"event_description" gorm:"not null"`
	Timestamp
}

type EventHubEventDTO struct {
	ID
	EventName        string `json:"event_name"`
	EventLocation    string `json:"event_location"`
	EventTime        string `json:"event_time"`
	EventDescription string `json:"event_description"`
	TimestampString
}

func (EventHubEvent) TableName() string {
	return tablePrefix + "users"
}
