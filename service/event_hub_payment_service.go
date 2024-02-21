package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubPaymentService = newEventHubPaymentService()

type eventHubPaymentService struct {
}

func newEventHubPaymentService() eventHubPaymentService {
	return eventHubPaymentService{}
}

func (s eventHubPaymentService) AddPaymentTransaction(configuration models.EventHubPaymentTransactions) error {
	/*---------------------------------------------------------
	 01. ADD CONFIGURATION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	_, dbResponse := repositories.EventHubPaymentRepository.AddPaymentTransaction(&configuration)
	if dbResponse.RowsAffected == 0 {
		return errors.New("failed to add payment transaction! ")
	}
	return nil
}

func (s eventHubPaymentService) GetPaymentTransactions(pagination models.Pagination, query, status string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	paymentTransactions, dbResponse := repositories.EventHubPaymentRepository.GetPaymentTransactions(pagination, newQuery, status)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("payment transaction not found! ")
	}
	return paymentTransactions, nil
}
