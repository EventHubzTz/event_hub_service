package models

type EventHubVotingPaymentTransactions struct {
	ID
	OrderID        string  `json:"order_id" gorm:"null"`
	TransactionID  string  `json:"transaction_id" gorm:"not null"`
	PhoneNumber    string  `json:"phone_number" gorm:"not null"`
	TotalAmount    float32 `json:"total_amount" gorm:"default:0;not null"`
	Currency       string  `json:"currency" gorm:"not null"`
	Provider       string  `json:"provider" gorm:"not null"`
	PaymentStatus  string  `json:"payment_status" gorm:"not null;type:enum('PENDING','COMPLETED','CANCELLED');default:'PENDING'"`
	GeneratedID    string  `json:"generated_id" gorm:"not null"`
	VotedFor       string  `json:"voted_for"  gorm:"not null"`
	VoteNumbers    int     `json:"vote_numbers" gorm:"not null"`
	VotedForCode   string  `json:"voted_for_code" gorm:"not null"`
	Longitude      string  `json:"longitude" gorm:"null"`
	Latitude       string  `json:"latitude" gorm:"null"`
	VotedID        string  `json:"voted_id" gorm:"not null"`
	Browser        string  `json:"browser" gorm:"null"`
	OS             string  `json:"os" gorm:"null"`
	UserAgent      string  `json:"user_agent" gorm:"null"`
	Device         string  `json:"device" gorm:"null"`
	OsVersion      string  `json:"os_version" gorm:"null"`
	BrowserVersion string  `json:"browser_version" gorm:"null"`
	DeviceType     string  `json:"device_type" gorm:"null"`
	IPAddress      string  `json:"ipaddress" gorm:"null"`
	Orientation    string  `json:"orientation" gorm:"null"`
	Location       string  `json:"location" gorm:"null"`
	Timestamp
}

/*----------------------------------------
  01.  DATA TRANSFER OBJECT
------------------------------------------*/

type EventHubVotingPaymentTransactionsDTO struct {
	ID             uint64  `json:"id"`
	OrderID        string  `json:"order_id"`
	TransactionID  string  `json:"transaction_id"`
	PhoneNumber    string  `json:"phone_number"`
	TotalAmount    float32 `json:"total_amount"`
	Currency       string  `json:"currency"`
	Provider       string  `json:"provider"`
	PaymentStatus  string  `json:"payment_status"`
	GeneratedID    string  `json:"generated_id"`
	VotedFor       string  `json:"voted_for"`
	VoteNumbers    int     `json:"vote_numbers"`
	VotedForCode   string  `json:"voted_for_code"`
	Longitude      string  `json:"longitude"`
	Latitude       string  `json:"latitude"`
	VotedID        string  `json:"voted_id"`
	Browser        string  `json:"browser"`
	OS             string  `json:"os"`
	UserAgent      string  `json:"user_agent"`
	Device         string  `json:"device"`
	OsVersion      string  `json:"os_version"`
	BrowserVersion string  `json:"browser_version"`
	DeviceType     string  `json:"device_type"`
	IPAddress      string  `json:"Ipaddress"`
	Orientation    string  `json:"orientation"`
	Location       string  `json:"location"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func (EventHubVotingPaymentTransactions) TableName() string {
	return tablePrefix + "voting_payment_transactions"
}
