package models

type EventHubRequestID struct {
	ID
	RequestID string `json:"request_id" gorm:"not null; unique"`
	Timestamp
}

/*
--------------------------------
 02. USER DATA TRANSFER OBJECT

---------------------------------
*/
type EventHubRequestIDDTO struct {
	Id        int64  `json:"ID"`
	RequestID string `json:"REQUEST_KEY"`
}

func (EventHubRequestID) TableName() string {
	return tablePrefix + "request_ids"
}
