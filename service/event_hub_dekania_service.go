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

func (s eventHubDekaniaService) AddRegion(region models.EventHubRegion) error {
	_, dbResp := repositories.EventHubDekaniaRepository.AddRegion(region)
	if dbResp.RowsAffected == 0 {
		return errors.New("fail to add region ")
	}
	return nil
}

func (s eventHubDekaniaService) GetAllRegionsByPagination(pagination models.Pagination, query string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	region, dbResponse := repositories.EventHubDekaniaRepository.GetAllRegionsByPagination(pagination, newQuery)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("regions not found! ")
	}
	return region, nil
}

func (s eventHubDekaniaService) UpdateRegion(regionRequest models.EventHubRegion, id uint64) error {
	/*--------------------------------------------------------------------
	 01. FIND REGION WITH GIVEN ID
	-----------------------------------------------------------------------*/
	region, dbResponse := repositories.EventHubDekaniaRepository.GetRegionWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update region! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE REGION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	region.RegionName = regionRequest.RegionName
	dbResponse = repositories.EventHubDekaniaRepository.UpdateRegionWithId(region)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update region! ")
	}

	return nil
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
	dekania, dbResponse := repositories.EventHubDekaniaRepository.GetDekaniaWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update dekania! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE REGION AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	dekania.DekaniaName = regionRequest.DekaniaName
	dbResponse = repositories.EventHubDekaniaRepository.UpdateDekaniaWithId(dekania)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update dekania! ")
	}

	return nil
}
