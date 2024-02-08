package models

type EventHubExternalOperationsSetup struct {
	ID
	Parameter string `json:"parameter" gorm:"not null;size:255"`
	Value     string `json:"value" gorm:"not null;type:text"`
	Timestamp
}

func (EventHubExternalOperationsSetup) TableName() string {
	return tablePrefix + "external_operations_setup"
}
