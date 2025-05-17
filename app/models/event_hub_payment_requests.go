package models

type EventHubPaymentRequests struct {
	ID
	FirstName     string  `json:"first_name" gorm:"not null"`
	LastName      string  `json:"last_name" gorm:"not null"`
	AccountNumber string  `json:"account_number" gorm:"not null"`
	BankName      string  `json:"bank_name" gorm:"null"`
	Amount        float32 `json:"amount" gorm:"default:0;not null"`
	PaymentStatus string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	Timestamp
}

type EventHubPaymentRequestsDTO struct {
	ID            uint64  `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber string  `json:"account_number"`
	BankName      string  `json:"bank_name"`
	Amount        float32 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (EventHubPaymentRequests) TableName() string {
	return tablePrefix + "payment_requests"
}
