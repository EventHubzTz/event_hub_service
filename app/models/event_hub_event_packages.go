package models

type EventHubEventPackages struct {
	ID
	EventID     uint64  `json:"event_id" gorm:"not null;index:event_id_index"`
	PackageName string  `json:"package_name" gorm:"not null"`
	Amount      float32 `json:"amount" gorm:"default:0;not null"`
	Timestamp

	EventHubEvent EventHubEvent `gorm:"foreignKey:EventID;constraint:OnDelete:NO ACTION"`
}

type EventHubEventPackagesDTO struct {
	ID          uint64  `json:"id"`
	EventID     uint64  `json:"-"`
	PackageName string  `json:"package_name"`
	Amount      float32 `json:"amount"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

/*----------------------------------------
  01.  EVENT PACKAGES ORDERS DATA TRANSFER OBJECT
------------------------------------------*/

func (EventHubEventPackages) TableName() string {
	return tablePrefix + "event_packages"
}
