package models

type EventHubVotingPaymentTransactions struct {
	ID
	OrderID       string  `json:"order_id" gorm:"null"`
	TransactionID string  `json:"transaction_id" gorm:"not null"`
	NomineeID     string  `json:"nominee_id" gorm:"not null"`
	NumberOfVotes int     `json:"number_of_votes" gorm:"default:1;not null"`
	PhoneNumber   string  `json:"phone_number" gorm:"not null"`
	Amount        float32 `json:"amount" gorm:"default:0;not null"`
	Currency      string  `json:"currency" gorm:"not null"`
	Provider      string  `json:"provider"  gorm:"not null"`
	PaymentStatus string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	Date          string  `json:"date"  gorm:"not null"`
	Timestamp
}

/*----------------------------------------
  01.  DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubVotingPaymentTransactionsDTO struct {
	ID            uint64  `json:"id"`
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	NomineeID     string  `json:"nominee_id"`
	NumberOfVotes int     `json:"number_of_votes"`
	PhoneNumber   string  `json:"phone_number"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	PaymentStatus string  `json:"payment_status"`
	Date          string  `json:"date"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (EventHubVotingPaymentTransactions) TableName() string {
	return tablePrefix + "voting_payment_transactions"
}
