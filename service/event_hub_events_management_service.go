package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubEventsManagementService = newEventHubEventsManagementService()

type eventHubEventsManagementService struct {
}

func newEventHubEventsManagementService() eventHubEventsManagementService {
	return eventHubEventsManagementService{}
}

func (s eventHubEventsManagementService) AddEvent(event models.EventHubEvent) error {
	/*---------------------------------------------------------
	 01. CREATE EVENT AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	_, dbResponse := repositories.EventHubEventsManagementRepository.AddEvent(&event)
	if dbResponse.RowsAffected == 0 {
		return errors.New("failed to register event! ")
	}
	return nil
}

func (s eventHubEventsManagementService) AddEventImage(eventImage models.EventHubEventImages) error {
	_, dbResp := repositories.EventHubEventsManagementRepository.AddEventImage(eventImage)
	if dbResp.RowsAffected == 0 {
		return errors.New("fail to add event image! ")
	}
	return nil
}

func (s eventHubEventsManagementService) GetEvents(pagination models.Pagination, query string, eventCategoryId, eventSubCategoryId uint64) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	events, dbResponse := repositories.EventHubEventsManagementRepository.GetEvents(pagination, newQuery, eventCategoryId, eventSubCategoryId)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("events not found! ")
	}
	return events, nil
}

func (s eventHubEventsManagementService) GetEvent(id uint64) (*models.EventHubEventDTO, error) {
	event, dbResponse := repositories.EventHubEventsManagementRepository.GetEvent(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return nil, dbResponse.Error
	}
	return &event, nil
}

func (s eventHubEventsManagementService) UpdateEvent(regionRequest models.EventHubEvent, id uint64) error {
	/*--------------------------------------------------------------------
	 01. FIND REGION WITH GIVEN ID
	-----------------------------------------------------------------------*/
	event, dbResponse := repositories.EventHubEventsManagementRepository.GetEventWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE REGION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	event.EventName = regionRequest.EventName
	event.EventLocation = regionRequest.EventLocation
	event.EventTime = regionRequest.EventTime
	event.EventDescription = regionRequest.EventDescription
	event.EventCapacity = regionRequest.EventCapacity
	event.EventEntrance = regionRequest.EventEntrance
	event.EventCategoryID = regionRequest.EventCategoryID
	event.EventSubCategoryID = regionRequest.EventSubCategoryID
	dbResponse = repositories.EventHubEventsManagementRepository.UpdateEventWithId(event)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event! ")
	}

	return nil
}

func (s eventHubEventsManagementService) CheckIfEventReachMaxCoverImageLimit(eventID uint64) error {
	/*---------------------------------------------------
	 01.  CHECKING IF PRODUCT REACH MAX IMAGE LIMIT (5)
	----------------------------------------------------*/
	coverImageFromDB, _ := repositories.EventHubEventsManagementRepository.FindProductImagesByProductID(eventID)
	if len(coverImageFromDB) >= 5 {
		return errors.New("you have reach  maximum number of 5 photo per event")
	}
	return nil
}

func (s eventHubEventsManagementService) GetDashboardStatistics() *models.EventHubDashboardStatisticsDTO {
	statistics, _ := repositories.EventHubEventsManagementRepository.GetDashboardStatistics()
	return statistics
}
