package models

import (
	"time"

	"gorm.io/gorm"
)

const tablePrefix = "event_hub_"

var Tables = []interface{}{
	&EventHubExternalOperationsSetup{}, &EventHubRequestID{}, &EventHubUser{},
	&EventHubUserToken{}, &EventHubUserOTPCode{}, &EventHubOTPCodeMessage{},
	&EventHubOTPMessageResponse{}, &EventHubForgotPasswordOTP{}, &EventHubEvent{},
	&EventHubEventCategories{}, &EventHubEventSubCategories{}, &EventHubEventImages{},
	&EventHubConfigurations{}, &EventHubPaymentTransactions{}, &EventHubEventPackages{},
	&EventHubDekania{}, &EventHubVotingPaymentTransactions{}, &EventHubRegion{},
	&EventHubContributionTransactions{},
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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ImageDimension struct {
	Width  int
	Height int
}
