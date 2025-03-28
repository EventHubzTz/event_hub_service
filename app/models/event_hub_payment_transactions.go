package models

type EventHubPaymentTransactions struct {
	ID
	OrderID              string  `json:"order_id" gorm:"null"`
	TransactionID        string  `json:"transaction_id" gorm:"not null"`
	EventID              uint64  `json:"event_id" gorm:"not null;index:event_id_index"`
	UserID               uint64  `json:"user_id" gorm:"not null;index:users_products_user_id_index"`
	TicketOwnerFirstName string  `json:"ticket_owner_first_name" gorm:"not null"`
	TicketOwnerLastName  string  `json:"ticket_owner_last_name" gorm:"not null"`
	TShirtSize           string  `json:"t_shirt_size" gorm:"null"`
	Region               string  `json:"region" gorm:"not null"`
	Location             string  `json:"location" gorm:"null"`
	Distance             string  `json:"distance" gorm:"null"`
	DateOfBirth          string  `json:"date_of_birth" gorm:"null;size:50"`
	PhoneNumber          string  `json:"phone_number" gorm:"not null"`
	Amount               float32 `json:"amount" gorm:"default:0;not null"`
	Currency             string  `json:"currency" gorm:"not null"`
	Provider             string  `json:"provider"  gorm:"not null"`
	PaymentStatus        string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	Timestamp

	EventHubUser  EventHubUser  `gorm:"foreignKey:UserID;constraint:OnDelete:NO ACTION"`
	EventHubEvent EventHubEvent `gorm:"foreignKey:EventID;constraint:OnDelete:NO ACTION"`
}

/*----------------------------------------
  01.  DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubPaymentTransactionsDTO struct {
	ID            uint64  `json:"id"`
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	EventID       uint64  `json:"event_id"`
	UserID        uint64  `json:"user_id"`
	EventName     string  `json:"event_name"`
	FullName      string  `json:"full_name"`
	TicketOwner   string  `json:"ticket_owner"`
	TShirtSize    string  `json:"t_shirt_size"`
	Region        string  `json:"region"`
	Location      string  `json:"location"`
	Distance      string  `json:"distance"`
	DateOfBirth   string  `json:"date_of_birth"`
	PhoneNumber   string  `json:"phone_number"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	PaymentStatus string  `json:"payment_status"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func (EventHubPaymentTransactions) TableName() string {
	return tablePrefix + "payment_transactions"
}
