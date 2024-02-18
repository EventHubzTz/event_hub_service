package models

type EventHubEventSubCategories struct {
	ID
	EventSubCategoryName string `json:"event_sub_category_name" gorm:"not null"`
	EventCategoryID      uint64 `json:"event_category_id" gorm:"not null;index:products_category_id_index"`
	IconUrl              string `json:"icon_url" gorm:"null;size:1000"`
	ImageStorage         string `json:"image_storage" gorm:"not null;type:enum('LOCAL','REMOTE');default:'LOCAL'"`
	Timestamp

	//FOREIGN KEY
	EventHubEventCategories EventHubEventCategories `gorm:"foreignKey:EventCategoryID;constraint:OnDelete:CASCADE"`
}

/*----------------------------------------
  01.  PRODUCTS SUB CATEGORIES DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubEventSubCategoriesDTO struct {
	ID                   uint64 `json:"id"`
	EventSubCategoryName string `json:"event_sub_category_name"`
	EventCategoryID      uint64 `json:"event_category_id"`
	IconUrl              string `json:"icon_url"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}

func (EventHubEventSubCategories) TableName() string {
	return tablePrefix + "event_subcategories"
}
