package models

import "time"

type EventHubUserToken struct {
	Id        uint64    `json:"id" gorm:"primarykey"`
	UserID    uint64    `json:"user_id" gorm:"not null;index:users_tokens_user_id_index"` //FOREIGN KEY
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at" gorm:"default:null"`
	Status    int       `json:"status" gorm:"not null;index:idx_user_token_status;default:1" form:"status"`
	Timestamp

	//FOREIGN KEYS
	EventHubUser EventHubUser `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (EventHubUserToken) TableName() string {
	return tablePrefix + "users_tokens"
}
