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
	Provider      string  `json:"provider"`
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
	}
}
