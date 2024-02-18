package categories_sub_categories

import "github.com/EventHubzTz/event_hub_service/app/models"

type EventHubEventCategoriesRequest struct {
	EventCategoryName  string `json:"event_category_name" validate:"required"`
	IconUrl            string `json:"icon_url"`
	ImageStorage       string `json:"image_storage"`
	EventCategoryColor string `json:"event_category_color" validate:"required"`
}

type EventHubEventCategoriesUpdateRequest struct {
	models.IDRequest
	EventCategoryName  string `json:"event_category_name" validate:"required"`
	IconUrl            string `json:"icon_url"`
	ImageStorage       string `json:"image_storage"`
	EventCategoryColor string `json:"event_category_color" validate:"required"`
}

type EventHubEventCategoriesGetRequest struct {
	EventCategoryID uint64 `json:"id" validate:"required"`
}

func (request EventHubEventCategoriesRequest) ToModel() models.EventHubEventCategories {
	return models.EventHubEventCategories{
		EventCategoryName:  request.EventCategoryName,
		IconUrl:            request.IconUrl,
		ImageStorage:       request.ImageStorage,
		EventCategoryColor: request.EventCategoryColor,
	}
}

func (request EventHubEventCategoriesUpdateRequest) ToModel() models.EventHubEventCategories {
	return models.EventHubEventCategories{
		EventCategoryName:  request.EventCategoryName,
		IconUrl:            request.IconUrl,
		ImageStorage:       request.ImageStorage,
		EventCategoryColor: request.EventCategoryColor,
	}
}

type EventHubEventSubCategoriesRequest struct {
	EventSubCategoryName string `json:"event_sub_category_name" validate:"required"`
	EventCategoryID      uint64 `json:"event_category_id" validate:"required"`
	IconUrl              string `json:"icon_url"`
	ImageStorage         string `json:"image_storage"`
}

type EventHubEventSubCategoriesUpdateRequest struct {
	models.IDRequest
	EventSubCategoryName string `json:"event_sub_category_name" validate:"required"`
	IconUrl              string `json:"icon_url"`
	ImageStorage         string `json:"image_storage"`
}

type SubCategoryPaginationRequest struct {
	Query           string `json:"query"`
	EventCategoryID uint64 `json:"event_category_id" validate:"required"`
	Limit           int    `json:"limit,omitempty" `
	Page            int    `json:"page,omitempty"`
	Sort            string `json:"sort,omitempty"`
}

type EventHubEventSubCategoriesGetRequest struct {
	EventSubCategoryID uint64 `json:"id" validate:"required"`
}

func (request EventHubEventSubCategoriesRequest) ToModel() models.EventHubEventSubCategories {
	return models.EventHubEventSubCategories{
		EventSubCategoryName: request.EventSubCategoryName,
		EventCategoryID:      request.EventCategoryID,
		IconUrl:              request.IconUrl,
		ImageStorage:         request.ImageStorage,
	}
}

func (request EventHubEventSubCategoriesUpdateRequest) ToModel() models.EventHubEventSubCategories {
	return models.EventHubEventSubCategories{
		EventSubCategoryName: request.EventSubCategoryName,
		IconUrl:              request.IconUrl,
		ImageStorage:         request.ImageStorage,
	}
}
