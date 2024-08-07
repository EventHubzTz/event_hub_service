package payments

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubGetTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
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
	Location             string  `json:"location" validate:"required"`
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
		Location:             request.Location,
		Distance:             request.Distance,
		DateOfBirth:          request.DateOfBirth,
		PhoneNumber:          request.PhoneNumber,
		Amount:               request.Amount,
		Currency:             request.Currency,
		Provider:             request.Provider,
	}
}

type EventHubVotingPaymentRequest struct {
	OrderID        string  `json:"order_id"`
	TransactionID  string  `json:"transaction_id"`
	PhoneNumber    string  `json:"phone_number" validate:"required"`
	TotalAmount    float32 `json:"amount"`
	Currency       string  `json:"currency"`
	Provider       string  `json:"provider" validate:"required"`
	GeneratedID    string  `json:"generated_id" validate:"required"`
	VotedFor       string  `json:"voted_for" validate:"required"`
	VoteNumbers    int     `json:"vote_numbers" validate:"required"`
	VotedForCode   string  `json:"voted_for_code" validate:"required"`
	Longitude      string  `json:"longitude"`
	Latitude       string  `json:"latitude"`
	VotedID        string  `json:"voted_id" validate:"required"`
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
}

func (request EventHubVotingPaymentRequest) ToModel() models.EventHubVotingPaymentTransactions {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubVotingPaymentTransactions{
		OrderID:        request.OrderID,
		TransactionID:  request.TransactionID,
		PhoneNumber:    request.PhoneNumber,
		TotalAmount:    request.TotalAmount,
		Currency:       request.Currency,
		Provider:       request.Provider,
		GeneratedID:    request.GeneratedID,
		VotedFor:       request.VotedFor,
		VoteNumbers:    request.VoteNumbers,
		VotedForCode:   request.VotedForCode,
		Longitude:      request.Longitude,
		Latitude:       request.Latitude,
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
