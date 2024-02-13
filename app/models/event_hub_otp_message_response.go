package models

type EventHubOTPMessageResponse struct {
	ID
	Value string `json:"value" gorm:"type:json"`
	Timestamp
}

func (EventHubOTPMessageResponse) TableName() string {
	return tablePrefix + "otp_message_response"
}
