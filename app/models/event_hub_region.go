package models

type EventHubRegion struct {
	ID
	RegionName string `json:"region_name" gorm:"not null"`
	Timestamp
}

/*----------------------------------------
  01.  DEKANIA DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubRegionDTO struct {
	ID         uint64 `json:"id"`
	RegionName string `json:"region_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (EventHubRegion) TableName() string {
	return tablePrefix + "regions"
}
