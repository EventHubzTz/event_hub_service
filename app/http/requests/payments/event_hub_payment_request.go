package payments

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubGetTransactionRequest struct {
	TransactionID string `json:"transaction_id"`
}

type EventHubPaymentRequest struct {
	OrderID        string  `json:"order_id"`
	TransactionID  string  `json:"transaction_id"`
	EventID        uint64  `json:"event_id" validate:"required"`
	EventPackageID uint64  `json:"event_package_id" validate:"required"`
	UserID         uint64  `json:"user_id"`
	TicketOwner    string  `json:"ticket_owner" validate:"required"`
	TShirtSize     string  `json:"t_shirt_size"`
	Location       string  `json:"location" validate:"required"`
	Distance       string  `json:"distance"`
	DateOfBirth    string  `json:"date_of_birth"`
	PhoneNumber    string  `json:"phone_number" validate:"required"`
	Amount         float32 `json:"amount"`
	Currency       string  `json:"currency"`
	Provider       string  `json:"provider" validate:"required"`
}

func (request EventHubPaymentRequest) ToModel() models.EventHubPaymentTransactions {
	/*---------------------------------------------------------
	 01. ASSIGN REQUEST TO EVENT MODEL
	----------------------------------------------------------*/
	return models.EventHubPaymentTransactions{
		OrderID:       request.OrderID,
		TransactionID: request.TransactionID,
		EventID:       request.EventID,
		UserID:        request.UserID,
		TicketOwner:   request.TicketOwner,
		TShirtSize:    request.TShirtSize,
		Location:      request.Location,
		Distance:      request.Distance,
		DateOfBirth:   request.DateOfBirth,
		PhoneNumber:   request.PhoneNumber,
		Amount:        request.Amount,
		Currency:      request.Currency,
		Provider:      request.Provider,
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
