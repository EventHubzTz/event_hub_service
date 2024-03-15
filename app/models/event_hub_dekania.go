package models

type EventHubDekania struct {
	ID
	DekaniaName string `json:"dekania_name" gorm:"not null"`
	Timestamp
}

/*----------------------------------------
  01.  DEKANIA DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubDekaniaDTO struct {
	ID          uint64 `json:"id"`
	DekaniaName string `json:"dekania_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (EventHubDekania) TableName() string {
	return tablePrefix + "dekania"
}
