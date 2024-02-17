package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/utils"
)

var EventHubRequestIDService = newEventHubRequestIDService()

type eventHubRequestIDService struct {
}

func newEventHubRequestIDService() eventHubRequestIDService {
	return eventHubRequestIDService{}
}

func (s eventHubRequestIDService) GetRequestID() *models.EventHubRequestIDDTO {
	requestID, usDB := repositories.EventHubRequestIDRepository.GetRequestID()
	if usDB.RowsAffected == 0 {
		return nil
	}
	return requestID
}

func (s eventHubRequestIDService) VerifyRequestID(requestID string) (bool, error) {
	request, usDB := repositories.EventHubRequestIDRepository.GetRequestID()
	if usDB.RowsAffected == 0 {
		return false, errors.New("request ID not found! ")
	}
	if request.RequestID == requestID {
		utils.SuccessPrint("OK")
		return true, nil
	}
	utils.InfoPrint("db Req ID: " + request.RequestID)
	utils.InfoPrint("sent Req ID: " + requestID)
	utils.ErrorPrint("Unauthhorized access")
	return false, errors.New("Unauthhorized")
}
