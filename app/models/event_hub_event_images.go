package models

type EventHubEventImages struct {
	ID
	EventID       uint64 `json:"event_id" gorm:"not null;index:products_id_index"`
	ImageUrl      string `json:"image_url" gorm:"column:image_url;type:longtext;not null"`
	ThumbunailUrl string `json:"thumbunail_url" gorm:"null;size:255"`
	VideoUrl      string `json:"video_url" gorm:"null;size:255"`
	ImageStorage  string `json:"image_storage" gorm:"not null;type:enum('LOCAL','REMOTE');default:'LOCAL'"`
	AspectRatios  string `json:"aspect_ratios" gorm:"null;size:10"`
	FileType      string `json:"file_type" gorm:"not null;type:enum('IMAGE','VIDEO','YOUTUBE');default:'IMAGE'"`
	Timestamp

	//FOREIGN KEYS
	EventHubEvent EventHubEvent `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}

/*
------------------------------------
 01. EVENT IMAGE DATA TRANSFER OBJECT
    -------------------------------------
*/
type EventHubEventImagesDTO struct {
	ID            uint64 `json:"id"`
	EventID       uint64 `json:"-"`
	ImageUrl      string `json:"image_url"`
	ThumbunailUrl string `json:"thumbunail_url"`
	VideoUrl      string `json:"video_url"`
	ImageStorage  string `json:"image_storage"`
	AspectRatios  string `json:"aspect_ratios"`
	FileType      string `json:"file_type"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (EventHubEventImages) TableName() string {
	return tablePrefix + "event_images"
}
