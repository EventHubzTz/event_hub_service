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

func (s eventHubPaymentService) AddPaymentTransaction(paymentData models.EventHubPaymentTransactions) error {
	/*---------------------------------------------------------
	 01. ADD CONFIGURATION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	_, dbResponse := repositories.EventHubPaymentRepository.AddPaymentTransaction(&paymentData)
	if dbResponse.RowsAffected == 0 {
		return errors.New("failed to add payment transaction! ")
	}
	return nil
}

func (s eventHubPaymentService) AddVotingPaymentTransaction(paymentData models.EventHubVotingPaymentTransactions) error {
	/*---------------------------------------------------------
	 01. ADD CONFIGURATION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	_, dbResponse := repositories.EventHubPaymentRepository.AddVotingPaymentTransaction(&paymentData)
	if dbResponse.RowsAffected == 0 {
		return errors.New("failed to add payment transaction! ")
	}
	return nil
}

func (s eventHubPaymentService) GetVotingPaymentTransactions(pagination models.Pagination, query, status string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	paymentTransactions, dbResponse := repositories.EventHubPaymentRepository.GetVotingPaymentTransactions(pagination, newQuery, status)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("payment transaction not found! ")
	}
	return paymentTransactions, nil
}

func (s eventHubPaymentService) GetPaymentTransactions(pagination models.Pagination, role, query, status, phoneNumber string, userID uint64) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	paymentTransactions, dbResponse := repositories.EventHubPaymentRepository.GetPaymentTransactions(pagination, role, newQuery, status, phoneNumber, userID)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("payment transaction not found! ")
	}
	return paymentTransactions, nil
}
