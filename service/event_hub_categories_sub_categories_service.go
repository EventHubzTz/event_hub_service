package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubCategoriesSubCategoriesService = newEventHubCategoriesSubCategoriesService()

type eventHubCategoriesSubCategoriesService struct {
}

func newEventHubCategoriesSubCategoriesService() eventHubCategoriesSubCategoriesService {
	return eventHubCategoriesSubCategoriesService{}
}

func (s eventHubCategoriesSubCategoriesService) CreateEventCategory(eventCategory models.EventHubEventCategories) error {
	_, dbResp := repositories.EventHubCategoriesSubCategoriesRepository.CreateEventCategory(eventCategory)
	if dbResp.RowsAffected == 0 {
		return errors.New("fail to add event category ")
	}
	return nil
}

func (s eventHubCategoriesSubCategoriesService) GetAllEventCategoriesByPagination(pagination models.Pagination, query string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	eventCategories, dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.GetAllEventCategoriesByPagination(pagination, newQuery)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("event categories not found! ")
	}
	return eventCategories, nil
}

func (s eventHubCategoriesSubCategoriesService) UpdateEventCategory(eventCategoryRequest models.EventHubEventCategories, id uint64) error {
	/*--------------------------------------------------------------------
	 01. FIND EVENT CATEGORY WITH GIVEN ID
	-----------------------------------------------------------------------*/
	eventCategory, dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.GetEventCategoryWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event category! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE EVENT CATEGORY AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	eventCategory.EventCategoryName = eventCategoryRequest.EventCategoryName
	eventCategory.IconUrl = eventCategoryRequest.IconUrl
	eventCategory.ImageStorage = eventCategoryRequest.ImageStorage
	eventCategory.EventCategoryColor = eventCategoryRequest.EventCategoryColor
	dbResponse = repositories.EventHubCategoriesSubCategoriesRepository.UpdateEventCategoryWithId(eventCategory)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event category! ")
	}

	return nil
}

func (s eventHubCategoriesSubCategoriesService) CreateEventSubCategory(eventCategory models.EventHubEventSubCategories) error {
	_, dbResp := repositories.EventHubCategoriesSubCategoriesRepository.CreateEventSubCategory(eventCategory)
	if dbResp.RowsAffected == 0 {
		return errors.New("fail to add event sub category ")
	}
	return nil
}

func (s eventHubCategoriesSubCategoriesService) GetAllEventSubCategoriesByPagination(pagination models.Pagination, eventCategoryId uint64, query string) (models.Pagination, error) {
	var newQuery = "%" + query + "%"
	productsSubCategories, dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.GetAllEventSubCategoriesByPagination(pagination, eventCategoryId, newQuery)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return models.Pagination{}, errors.New("event sub categories not found! ")
	}
	return productsSubCategories, nil
}

func (s eventHubCategoriesSubCategoriesService) UpdateEventSubCategory(eventSubCategoryRequest models.EventHubEventSubCategories, id uint64) error {
	/*--------------------------------------------------------------------
	 01. FIND EVENT CATEGORY WITH GIVEN ID
	-----------------------------------------------------------------------*/
	eventSubCategory, dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.GetEventSubCategoryWithId(id)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event sub category! ")
	}
	/*--------------------------------------------------------------------
	 02. UPDATE EVENT CATEGORY AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	-----------------------------------------------------------------------*/
	eventSubCategory.EventSubCategoryName = eventSubCategoryRequest.EventSubCategoryName
	eventSubCategory.IconUrl = eventSubCategoryRequest.IconUrl
	eventSubCategory.ImageStorage = eventSubCategoryRequest.ImageStorage
	dbResponse = repositories.EventHubCategoriesSubCategoriesRepository.UpdateEventSubCategoryWithId(eventSubCategory)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return errors.New("failed to update event sub category! ")
	}

	return nil
}
