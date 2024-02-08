package models

import (
	"time"

	"gorm.io/gorm"
)

const tablePrefix = "event_hub_"

var Tables = []interface{}{&EventHubExternalOperationsSetup{},
	&EventHubRequestID{}, &EventHubMessageResponse{},
	&EventHubUser{}, &EventHubUserToken{},
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type TimestampWithDeletedAt struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ID struct {
	Id uint64 `json:"id" gorm:"primarykey;index:users_id_index"`
}

type IDRequest struct {
	Id uint64 `json:"id" validate:"required"`
}

type TimestampString struct {
	CreatedAt string `json:"CREATED_AT"`
	UpdatedAt string `json:"UPDATED_AT"`
}
