package models

import "time"

type EventHubEvent struct {
	ID
	UserID             uint64    `json:"user_id" gorm:"not null;index:users_events_user_id_index"`
	EventName          string    `json:"event_name" gorm:"not null"`
	EventLocation      string    `json:"event_location" gorm:"not null"`
	EventTime          time.Time `json:"event_time" gorm:"not null"`
	EventCapacity      int       `json:"event_capacity" gorm:"default:0;not null"`
	EventEntrance      float32   `json:"event_entrance" gorm:"default:0;not null"`
	EventDescription   string    `json:"event_description" gorm:"not null"`
	EventCategoryID    uint64    `json:"event_category_id" gorm:"not null;index:event_category_id_index"`
	EventSubCategoryID uint64    `json:"event_sub_category_id" gorm:"not null;index:event_sub_category_id_index"`
	Timestamp

	//FOREIGN KEY
	EventHubUser               EventHubUser               `gorm:"foreignKey:UserID;constraint:OnDelete:NO ACTION"`
	EventHubEventCategories    EventHubEventCategories    `gorm:"foreignKey:EventCategoryID;constraint:OnDelete:CASCADE"`
	EventHubEventSubCategories EventHubEventSubCategories `gorm:"foreignKey:EventSubCategoryID;constraint:OnDelete:CASCADE"`
}

type EventHubEventDTO struct {
	ID
	UserID               uint64                   `json:"user_id"`
	EventOwner           string                   `json:"event_owner"`
	EventName            string                   `json:"event_name"`
	EventLocation        string                   `json:"event_location"`
	EventTime            string                   `json:"event_time"`
	EventCapacity        int                      `json:"event_capacity"`
	EventEntrance        float32                  `json:"event_entrance"`
	EventDescription     string                   `json:"event_description"`
	EventCategoryID      uint64                   `json:"event_category_id"`
	EventSubCategoryID   uint64                   `json:"event_sub_category_id"`
	EventCategoryName    string                   `json:"event_category_name"`
	EventSubCategoryName string                   `json:"event_sub_category_name"`
	EventFiles           []EventHubEventImagesDTO `json:"event_files" gorm:"foreignKey:event_id"`
	TimestampString
}

type EventHubDashboardStatisticsDTO struct {
	TotalUsers  int `json:"total_users"`
	TotalEvents int `json:"total_events"`
}

func (EventHubEvent) TableName() string {
	return tablePrefix + "events"
}
