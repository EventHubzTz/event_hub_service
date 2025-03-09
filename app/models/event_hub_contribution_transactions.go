package models

type EventHubContributionTransactions struct {
	ID
	OrderID       string  `json:"order_id" gorm:"null"`
	TransactionID string  `json:"transaction_id" gorm:"not null"`
	FirstName     string  `json:"first_name" gorm:"not null"`
	LastName      string  `json:"last_name" gorm:"not null"`
	Region        string  `json:"region" gorm:"not null"`
	Location      string  `json:"location" gorm:"null"`
	PhoneNumber   string  `json:"phone_number" gorm:"not null"`
	Amount        float32 `json:"amount" gorm:"default:0;not null"`
	Currency      string  `json:"currency" gorm:"not null"`
	Provider      string  `json:"provider"  gorm:"not null"`
	PaymentStatus string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	Timestamp
}

/*----------------------------------------
  01.  DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubContributionTransactionsDTO struct {
	ID            uint64  `json:"id"`
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	FullName      string  `json:"full_name"`
	Region        string  `json:"region"`
	Location      string  `json:"location"`
	PhoneNumber   string  `json:"phone_number"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	PaymentStatus string  `json:"payment_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (EventHubContributionTransactions) TableName() string {
	return tablePrefix + "contribution_transactions"
}
