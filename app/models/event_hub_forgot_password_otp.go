package models

type EventHubForgotPasswordOTP struct {
	ID
	UserID    uint64 `json:"user_id" gorm:"not null;index:forgot_password_otp_user_id_index"` //FOREIGN KEY
	OTP       string `json:"otp" gorm:"not null;size:4"`
	Message   string `json:"message" gorm:"not null;size:250"`
	Phone     string `json:"phone" gorm:"not null;size:20" validate:"required,min=9"`
	IsOTPSent string `json:"is_otp_sent" gorm:"not null;type:enum('YES','NO');default:'NO'"`
	Timestamp

	//FOREIGN KEYS
	EventHubUser EventHubUser `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

/*
--------------------------------------------
 01. FORGOT PASSWORD DATA TRANSFER OBJECT

----------------------------------------------
*/
type AFYAAPPForgotPasswordOTPDTO struct {
	Id        uint64 `json:"ID"`
	UserID    uint64 `json:"USER_ID"`
	OTP       string `json:"OTP"`
	Message   string `json:"MESSAGE"`
	Phone     string `json:"PHONE_NUMBER"`
	IsOTPSent string `json:"IS_OTP_SENT"`
	CreatedAt string `json:"CREATED_AT"`
}

func (EventHubForgotPasswordOTP) TableName() string {
	return tablePrefix + "forgot_password_otp"
}
