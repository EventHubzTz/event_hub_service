package payments

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
)

type EventHubPaymentRequest struct {
	OrderID       string  `json:"order_id"`
	TransactionID string  `json:"transaction_id"`
	EventID       uint64  `json:"event_id" validate:"required"`
	UserID        uint64  `json:"user_id"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	Amount        float32 `json:"amount"`
	Currency      string  `json:"currency"`
	Provider      string  `json:"provider" validate:"required"`
	PaymentStatus string  `json:"payment_status"`
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
		PhoneNumber:   request.PhoneNumber,
		Amount:        request.Amount,
		Currency:      request.Currency,
		Provider:      request.Provider,
		PaymentStatus: request.PaymentStatus,
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
	SubmerchantAcc    string `json:"submerchantAcc" validate:"required"`
	FspReferenceId    string `json:"fspReferenceId"`
}
