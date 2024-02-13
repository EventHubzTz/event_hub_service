package models

type EventHubOTPCodeMessage struct {
	ID
	UserID      uint64 `json:"user_id" gorm:"not null;index:otp_code_messages_user_id_index"` //FOREIGN KEY
	Body        string `json:"body" gorm:"not null;size:250"`
	MessageSent string `json:"message_sent" gorm:"not null;type:enum('YES','NO');default:'NO'"`
	Timestamp

	//FOREIGN KEYS
	EventHubUser EventHubUser `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

/*
----------------------------------------
01.  OTP CODE MESSAGE DATA TRANSFER OBJECT
------------------------------------------
*/
type EventHubOTPCodeMessageDTO struct {
	ID          uint64 `json:"ID"`
	UserID      uint64 `json:"USER_ID"`
	Body        string `json:"BODY"`
	MessageSent string `json:"MESSAGE_SENT"`
	CreatedAt   string `json:"CREATED_AT"`
}

func (EventHubOTPCodeMessage) TableName() string {
	return tablePrefix + "otp_code_messages"
}
