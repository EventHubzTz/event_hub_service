package models

type EventHubPaymentTransactions struct {
	ID
	OrderID       string  `json:"order_id" gorm:"null"`
	TransactionID string  `json:"transaction_id" gorm:"not null"`
	EventID       uint64  `json:"event_id" gorm:"not null;index:event_id_index"`
	UserID        uint64  `json:"user_id" gorm:"not null;index:users_products_user_id_index"`
	PhoneNumber   string  `json:"phone_number" gorm:"not null"`
	Amount        float32 `json:"amount" gorm:"default:0;not null"`
	Currency      string  `json:"currency" gorm:"not null"`
	Provider      string  `json:"provider"  gorm:"not null"`
	PaymentStatus string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	Timestamp

	EventHubUser  EventHubUser  `gorm:"foreignKey:UserID;constraint:OnDelete:NO ACTION"`
	EventHubEvent EventHubEvent `gorm:"foreignKey:EventID;constraint:OnDelete:NO ACTION"`
}

type EventHubPaymentTransactionsDTO struct {
	ID            uint64  `json:"id"`
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	EventID       uint64  `json:"event_id"`
	UserID        uint64  `json:"user_id"`
	EventName     string  `json:"event_name"`
	FullName      string  `json:"full_name"`
	PhoneNumber   string  `json:"phone_number"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	PaymentStatus string  `json:"payment_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

/*----------------------------------------
  01.  PRODUCTS ORDERS DATA TRANSFER OBJECT
------------------------------------------*/

func (EventHubPaymentTransactions) TableName() string {
	return tablePrefix + "payment_transactions"
}
