package models

type EventHubEventCategories struct {
	ID
	EventCategoryName  string `json:"event_category_name" gorm:"not null"`
	IconUrl            string `json:"icon_url" gorm:"null;size:1000"`
	ImageStorage       string `json:"image_storage" gorm:"not null;type:enum('LOCAL','REMOTE');default:'LOCAL'"`
	EventCategoryColor string `json:"event_category_color" gorm:"not null"`
	Timestamp
}

/*----------------------------------------
  01.  EVENT CATEGORIES DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubEventCategoriesDTO struct {
	ID                 uint64 `json:"id"`
	EventCategoryName  string `json:"event_category_name"`
	IconUrl            string `json:"icon_url"`
	EventCategoryColor string `json:"event_category_color"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

func (EventHubEventCategories) TableName() string {
	return tablePrefix + "event_categories"
}
