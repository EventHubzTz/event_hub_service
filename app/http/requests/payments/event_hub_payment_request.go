package payments

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubGetTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
}

type EventHubAzamPayPaymentRequest struct {
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	TotalAmount   float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider" validate:"required"`
}

type EventHubPaymentRequest struct {
	OrderID              string  `json:"order_id"`
	TransactionID        string  `json:"transaction_id"`
	EventID              uint64  `json:"event_id" validate:"required"`
	EventPackageID       uint64  `json:"event_package_id" validate:"required"`
	UserID               uint64  `json:"user_id"`
	TicketOwnerFirstName string  `json:"ticket_owner_first_name" validate:"required"`
	TicketOwnerLastName  string  `json:"ticket_owner_last_name" validate:"required"`
	TShirtSize           string  `json:"t_shirt_size"`
	Region               string  `json:"region" validate:"required"`
	Location             string  `json:"location"`
	Distance             string  `json:"distance"`
	DateOfBirth          string  `json:"date_of_birth"`
	PhoneNumber          string  `json:"phone_number" validate:"required"`
	Amount               float32 `json:"amount"`
	Currency             string  `json:"currency"`
	Provider             string  `json:"provider" validate:"required"`
}

func (request EventHubPaymentRequest) ToModel() models.EventHubPaymentTransactions {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubPaymentTransactions{
		OrderID:              request.OrderID,
		TransactionID:        request.TransactionID,
		EventID:              request.EventID,
		UserID:               request.UserID,
		TicketOwnerFirstName: request.TicketOwnerFirstName,
		TicketOwnerLastName:  request.TicketOwnerLastName,
		TShirtSize:           request.TShirtSize,
		Region:               request.Region,
		Location:             request.Location,
		Distance:             request.Distance,
		DateOfBirth:          request.DateOfBirth,
		PhoneNumber:          request.PhoneNumber,
		Amount:               request.Amount,
		Currency:             request.Currency,
		Provider:             request.Provider,
	}
}

type EventHubContributionPaymentRequest struct {
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	FirstName     string  `json:"first_name" validate:"required"`
	LastName      string  `json:"last_name" validate:"required"`
	Region        string  `json:"region" validate:"required"`
	Location      string  `json:"location"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider" validate:"required"`
}

func (request EventHubContributionPaymentRequest) ToModel() models.EventHubContributionTransactions {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubContributionTransactions{
		OrderID:       request.OrderID,
		TransactionID: request.TransactionID,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Region:        request.Region,
		Location:      request.Location,
		PhoneNumber:   request.PhoneNumber,
		Amount:        request.Amount,
		Currency:      request.Currency,
		Provider:      request.Provider,
	}
}

type EventHubDebitRequest struct {
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	FirstName     string  `json:"first_name" validate:"required"`
	LastName      string  `json:"last_name" validate:"required"`
	Region        string  `json:"region" validate:"required"`
	Location      string  `json:"location"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider" validate:"required"`
	PaymentStatus string  `json:"payment_status"`
}

func (request EventHubDebitRequest) ToModel() models.EventHubDebits {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubDebits{
		OrderID:       request.OrderID,
		TransactionID: request.TransactionID,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Region:        request.Region,
		Location:      request.Location,
		PhoneNumber:   request.PhoneNumber,
		Amount:        request.Amount,
		Currency:      request.Currency,
		Provider:      request.Provider,
		PaymentStatus: request.PaymentStatus,
	}
}

type EventHubVotingPaymentRequest struct {
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	TotalAmount   float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider"`
	GeneratedID   string  `json:"generated_id" validate:"required"`
	Category      string  `json:"category"`
	VotedFor      string  `json:"voted_for" validate:"required"`
	VoteNumbers   int     `json:"vote_numbers" validate:"required"`
	VotedForCode  string  `json:"voted_for_code" validate:"required"`
	// Longitude      string  `json:"longitude"`
	// Latitude       string  `json:"latitude"`
	VotedID        string `json:"voted_id" validate:"required"`
	Browser        string `json:"browser"`
	OS             string `json:"os"`
	UserAgent      string `json:"user_agent"`
	Device         string `json:"device"`
	OsVersion      string `json:"os_version"`
	BrowserVersion string `json:"browser_version"`
	DeviceType     string `json:"device_type"`
	IPAddress      string `json:"Ipaddress"`
	Orientation    string `json:"orientation"`
	Location       string `json:"location"`
}

func (request EventHubVotingPaymentRequest) ToModel() models.EventHubVotingPaymentTransactions {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubVotingPaymentTransactions{
		OrderID:       request.OrderID,
		TransactionID: request.TransactionID,
		PhoneNumber:   request.PhoneNumber,
		TotalAmount:   request.TotalAmount,
		Currency:      request.Currency,
		Provider:      request.Provider,
		GeneratedID:   request.GeneratedID,
		Category:      request.Category,
		VotedFor:      request.VotedFor,
		VoteNumbers:   request.VoteNumbers,
		VotedForCode:  request.VotedForCode,
		// Longitude:      request.Longitude,
		// Latitude:       request.Latitude,
		VotedID:        request.VotedID,
		Browser:        request.Browser,
		OS:             request.OS,
		UserAgent:      request.UserAgent,
		Device:         request.Device,
		OsVersion:      request.OsVersion,
		BrowserVersion: request.BrowserVersion,
		DeviceType:     request.DeviceType,
		IPAddress:      request.IPAddress,
		Orientation:    request.Orientation,
		Location:       request.Location,
	}
}

type EventHubUpdatePaymentStatusRequest struct {
	AdditionalProperties struct {
		Property1 interface{} `json:"property1"`
		Property2 interface{} `json:"property2"`
	} `json:"additionalProperties"`
	Msisdn            string `json:"msisdn" validate:"required"`
	Amount            string `json:"amount" validate:"required"`
	Message           string `json:"message" validate:"required"`
	Utilityref        string `json:"utilityref" validate:"required"`
	Operator          string `json:"operator" validate:"required"`
	Reference         string `json:"reference" validate:"required"`
	Transactionstatus string `json:"transactionstatus" validate:"required"`
	SubmerchantAcc    string `json:"submerchantAcc"`
	FspReferenceId    string `json:"fspReferenceId"`
}

type EventHubRequestPaymentRequest struct {
	FirstName     string  `json:"first_name" validate:"required"`
	LastName      string  `json:"last_name" validate:"required"`
	AccountNumber string  `json:"account_number" validate:"required"`
	BankName      string  `json:"bank_name"`
	Amount        float32 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
}

func (request EventHubRequestPaymentRequest) ToModel() models.EventHubPaymentRequests {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubPaymentRequests{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		AccountNumber: request.AccountNumber,
		BankName:      request.BankName,
		Amount:        request.Amount,
		PaymentStatus: request.PaymentStatus,
	}
}

type EventHubOtherPaymentRequest struct {
	TransactionID string  `json:"transaction_id"`
	FullName      string  `json:"full_name" alidate:"required"`
	TShirtSize    string  `json:"t_shirt_size"`
	RegionName    string  `json:"region_name" validate:"required"`
	Location      string  `json:"location"`
	Distance      string  `json:"distance"`
	Age           string  `json:"age"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	Amount        float32 `json:"amount"`
}

func (request EventHubOtherPaymentRequest) ToModel() models.EventHubOtherPayments {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO MODEL
	----------------------------------------------------------*/
	return models.EventHubOtherPayments{
		TransactionID: request.TransactionID,
		FullName:      request.FullName,
		TShirtSize:    request.TShirtSize,
		RegionName:    request.RegionName,
		Location:      request.Location,
		Distance:      request.Distance,
		Age:           request.Age,
		PhoneNumber:   request.PhoneNumber,
		Amount:        request.Amount,
	}
}
