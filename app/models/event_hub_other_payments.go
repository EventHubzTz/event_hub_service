package models

type EventHubOtherPayments struct {
	ID
	TransactionID string  `json:"transaction_id" gorm:"not null"`
	FullName      string  `json:"full_name" gorm:"not null"`
	TShirtSize    string  `json:"t_shirt_size" gorm:"null"`
	Region        string  `json:"region" gorm:"not null"`
	Location      string  `json:"location" gorm:"null"`
	Distance      string  `json:"distance" gorm:"null"`
	Age           string  `json:"age" gorm:"null;size:50"`
	PhoneNumber   string  `json:"phone_number" gorm:"not null"`
	Amount        float32 `json:"amount" gorm:"default:0;not null"`
	Timestamp
}

/*----------------------------------------
  01.  DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubOtherPaymentsDTO struct {
	ID            uint64  `json:"id"`
	TransactionID string  `json:"transaction_id"`
	FullName      string  `json:"full_name"`
	TShirtSize    string  `json:"t_shirt_size"`
	Region        string  `json:"region"`
	Location      string  `json:"location"`
	Distance      string  `json:"distance"`
	Age           string  `json:"age"`
	PhoneNumber   string  `json:"phone_number"`
	Amount        float32 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (EventHubOtherPayments) TableName() string {
	return tablePrefix + "other_payments"
}
