package models

type EventHubUserOTPCode struct {
	ID
	Phone   string `json:"phone" gorm:"not null;size:20" validate:"required,min=9"`
	OTP     string `json:"otp" gorm:"not null;size:6"`
	Expired bool   `json:"expired" gorm:"not null;default:false"`
	Timestamp
}

func (EventHubUserOTPCode) TableName() string {
	return tablePrefix + "users_otp_codes"
}
