package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubDekaniaService = newEventHubDekaniaService()

type eventHubDekaniaService struct {
}

func newEventHubDekaniaService() eventHubDekaniaService {
	return eventHubDekaniaService{}
}

func (s eventHubDekaniaService) AddDekania(dekania models.EventHubDekania) error {
	_, dbResp := repositories.EventHubDekaniaRepository.AddDekania(dekania)
	if dbResp.RowsAffected == 0 {
		return errors.New("fail to add dekania ")
	}
	return nil
}

func (s eventHubDekaniaService) GetAllDekaniaByPagination(pagination models.Pagination, query string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	dekania, dbResponse := repositories.EventHubDekaniaRepository.GetAllDekaniaByPagination(pagination, newQuery)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("dekania not found! ")
	}
	return dekania, nil
}

func (s eventHubDekaniaService) UpdateDekania(regionRequest models.EventHubDekania, id uint64) error {
	/*--------------------------------------------------------------------
	 01. FIND REGION WITH GIVEN ID
	-----------------------------------------------------------------------*/
	dekania, dbResponse := repositories.EventHubDekaniaRepository.GetRegionWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update dekania! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE REGION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	dekania.DekaniaName = regionRequest.DekaniaName
	dbResponse = repositories.EventHubDekaniaRepository.UpdateRegionWithId(dekania)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update dekania! ")
	}

	return nil
}
