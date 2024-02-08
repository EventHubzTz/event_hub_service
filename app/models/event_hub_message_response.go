package models

type EventHubMessageResponse struct {
	ID
	Value string `json:"value" gorm:"type:json"`
	Timestamp
}

func (EventHubMessageResponse) TableName() string {
	return tablePrefix + "message_response"
}
