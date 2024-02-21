package repositories

import (
	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubPaymentRepository = newEventHubPaymentRepository()

type eventHubPaymentRepository struct {
}

func newEventHubPaymentRepository() eventHubPaymentRepository {
	return eventHubPaymentRepository{}
}

func (r eventHubPaymentRepository) AddPaymentTransaction(paymentTransation *models.EventHubPaymentTransactions) (*models.EventHubPaymentTransactions, *gorm.DB) {
	urDB := db.Create(&paymentTransation)
	return paymentTransation, urDB
}

func (r eventHubPaymentRepository) GetPaymentTransactions(pagination models.Pagination, query, status string) (models.Pagination, *gorm.DB) {

	events, urDB := helpers.EventHubQueryBuilder.QueryPaymentTransactions(pagination, query, status)

	return events, urDB
}

func (r eventHubPaymentRepository) UpdatePaymentStatus(transactionID string, status string) *gorm.DB {

	urDB := db.Model(models.EventHubPaymentTransactions{}).Where("transaction_id = ? ", transactionID).Update("payment_status", status)
	return urDB
}
