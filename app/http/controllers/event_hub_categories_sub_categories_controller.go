package controllers

import (
	"os"
	"strconv"
	"strings"

	"github.com/EventHubzTz/event_hub_service/app/http/requests/categories_sub_categories"
	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/random_string_generators"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
)

var EventHubCategoriesSubCategoriesController = newEventHubCategoriesSubCategoriesController()

type eventHubCategoriesSubCategoriesController struct {
}

func newEventHubCategoriesSubCategoriesController() eventHubCategoriesSubCategoriesController {
	return eventHubCategoriesSubCategoriesController{}
}

func (c eventHubCategoriesSubCategoriesController) CreateEventCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventCategoriesRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	request.EventCategoryName = ctx.FormValue("event_category_name")
	request.EventCategoryColor = ctx.FormValue("event_category_color")
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*------------------------------------------------------------
	  04. GET UPLOADED FILE
	 --------------------------------------------------------------*/
	file, fileErr := ctx.FormFile("image")

	if fileErr != nil {
		return response.ErrorResponse(fileErr.Error(), fiber.StatusBadRequest, ctx)
	}

	/*------------------------------------------------------------
	  05. VALIDATING UPLOADED FILE TYPE
	 --------------------------------------------------------------*/
	allowedExtensions := []string{"png", "jpg", "jpeg", "svg"}
	uploadedFileExtension := strings.ToLower(strings.Split(file.Filename, ".")[1])
	if uploadedFileExtension != allowedExtensions[0] && uploadedFileExtension != allowedExtensions[1] &&
		uploadedFileExtension != allowedExtensions[2] {
		return response.ErrorResponse("Invalid file format. Supported. 'png', 'jpg', 'jpeg', 'svg'", fiber.StatusBadRequest, ctx)
	}

	/*-----------------------------------------------
	  06. SAVING FILE TO THE LOCAL STORAGE LOCATION
	 ------------------------------------------------*/
	os.MkdirAll("public/event/images/", os.ModePerm)
	savedPath := "/event/images/" + strings.ToLower(random_string_generators.RandomString(20)) + "." + uploadedFileExtension
	pathData := "public" + savedPath

	saveErr := ctx.SaveFile(file, pathData)

	if saveErr != nil {
		return response.ErrorResponse(saveErr.Error(), fiber.StatusConflict, ctx)
	}

	/*------------------------------------------------------------
	  07. ADD SAVED PATH TO THE REQUEST OBJECT AS URL
	 --------------------------------------------------------------*/
	request.IconUrl = savedPath
	request.ImageStorage = "LOCAL"
	/*--------------------------------------------------------------------
	 08. CREATE EVENT CATEGORY
	-----------------------------------------------------------------------*/
	err := service.EventHubCategoriesSubCategoriesService.CreateEventCategory(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Event category added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) GetAllEventCategories(ctx *fiber.Ctx) error {

	productsCategories, dbErr := repositories.EventHubCategoriesSubCategoriesRepository.GetAllEventCategories()

	if dbErr.RowsAffected == 0 {
		return response.ErrorResponse("No records found in event categories database", fiber.StatusOK, ctx)
	}

	return response.InternalServiceDataResponse(productsCategories, fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) GetAllEventCategoriesByPagination(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request events.EventHubEventsGetsRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	var pagination models.Pagination
	pagination.Limit = request.Limit
	pagination.Sort = request.Sort
	pagination.Page = request.Page

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. GET EVENT CATEGORIES AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	productCategoriesData, err := service.EventHubCategoriesSubCategoriesService.GetAllEventCategoriesByPagination(pagination, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(productCategoriesData, fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) UpdateEventCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventCategoriesUpdateRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	eventCategoryID, conversionError := strconv.Atoi(ctx.FormValue("id"))
	if conversionError == nil {
		request.Id = uint64(eventCategoryID)
	}
	request.EventCategoryName = ctx.FormValue("event_category_name")
	request.EventCategoryColor = ctx.FormValue("event_category_color")
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.ValidateForUpdate(request, request.Id)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*------------------------------------------------------------
	  04. GET UPLOADED FILE
	 --------------------------------------------------------------*/
	file, fileErr := ctx.FormFile("image")
	savedPath := ""
	if file != nil {

		if fileErr != nil {
			return response.ErrorResponse(fileErr.Error(), fiber.StatusBadRequest, ctx)
		}

		/*------------------------------------------------------------
		  05. VALIDATING UPLOADED FILE TYPE
		 --------------------------------------------------------------*/
		allowedExtensions := []string{"png", "jpg", "jpeg", "svg"}
		uploadedFileExtension := strings.ToLower(strings.Split(file.Filename, ".")[1])
		if uploadedFileExtension != allowedExtensions[0] && uploadedFileExtension != allowedExtensions[1] &&
			uploadedFileExtension != allowedExtensions[2] {
			return response.ErrorResponse("Invalid file format. Supported. 'png', 'jpg', 'jpeg', 'svg'", fiber.StatusBadRequest, ctx)
		}

		/*-----------------------------------------------
		  06. SAVING FILE TO THE LOCAL STORAGE LOCATION
		 ------------------------------------------------*/
		os.MkdirAll("public/event/images/", os.ModePerm)
		savedPath = "/event/images/" + strings.ToLower(random_string_generators.RandomString(20)) + "." + uploadedFileExtension
		pathData := "public" + savedPath

		saveErr := ctx.SaveFile(file, pathData)

		if saveErr != nil {
			return response.ErrorResponse(saveErr.Error(), fiber.StatusConflict, ctx)
		}

	}
	/*------------------------------------------------------------
	  07. ADD SAVED PATH TO THE REQUEST OBJECT AS URL
	 --------------------------------------------------------------*/
	if savedPath != "" {
		request.IconUrl = savedPath
		request.ImageStorage = "LOCAL"
	} else {
		request.IconUrl = ctx.FormValue("image")
		request.ImageStorage = "REMOTE"
	}
	/*-----------------------------------------------------------------
	 04. UPDATE EVENT CATEGORY NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err := service.EventHubCategoriesSubCategoriesService.UpdateEventCategory(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event category name updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) DeleteEventCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventCategoriesGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 05. DELETE EVENT CATEGORY AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.DeleteEventCategory(request.EventCategoryID)
	/*---------------------------------------------------------
	 06. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete product category", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event category deleted successfully!", fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) CreateEventSubCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventSubCategoriesRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	request.EventSubCategoryName = ctx.FormValue("event_sub_category_name")
	eventCategoryID, conversionError := strconv.Atoi(ctx.FormValue("event_category_id"))
	if conversionError == nil {
		request.EventCategoryID = uint64(eventCategoryID)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*------------------------------------------------------------
	  04. GET UPLOADED FILE
	 --------------------------------------------------------------*/
	file, fileErr := ctx.FormFile("image")

	if fileErr != nil {
		return response.ErrorResponse(fileErr.Error(), fiber.StatusBadRequest, ctx)
	}

	/*------------------------------------------------------------
	  05. VALIDATING UPLOADED FILE TYPE
	 --------------------------------------------------------------*/
	allowedExtensions := []string{"png", "jpg", "jpeg"}
	uploadedFileExtension := strings.ToLower(strings.Split(file.Filename, ".")[1])
	if uploadedFileExtension != allowedExtensions[0] && uploadedFileExtension != allowedExtensions[1] &&
		uploadedFileExtension != allowedExtensions[2] {
		return response.ErrorResponse("Invalid file format. Supported. 'png', 'jpg', 'jpeg'", fiber.StatusBadRequest, ctx)
	}

	/*-----------------------------------------------
	  06. SAVING FILE TO THE LOCAL STORAGE LOCATION
	 ------------------------------------------------*/
	os.MkdirAll("public/products_subcategories/images/", os.ModePerm)
	savedPath := "/products_subcategories/images/" + strings.ToLower(random_string_generators.RandomString(20)) + "." + uploadedFileExtension
	pathData := "public" + savedPath

	saveErr := ctx.SaveFile(file, pathData)

	if saveErr != nil {
		return response.ErrorResponse(saveErr.Error(), fiber.StatusConflict, ctx)
	}

	/*------------------------------------------------------------
	  07. ADD SAVED PATH TO THE REQUEST OBJECT AS URL
	 --------------------------------------------------------------*/
	request.IconUrl = savedPath
	request.ImageStorage = "LOCAL"
	/*--------------------------------------------------------------------
	 08. CREATE EVENT SUB CATEGORY
	-----------------------------------------------------------------------*/
	err := service.EventHubCategoriesSubCategoriesService.CreateEventSubCategory(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Event sub category added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) GetAllEventSubCategories(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventCategoriesGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. GET EVENT SUB CATEGORIES AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	productsSubCategories, dbErr := repositories.EventHubCategoriesSubCategoriesRepository.GetAllEventSubCategories(request.EventCategoryID)
	if dbErr.RowsAffected == 0 {
		return response.ErrorResponse("No records found in event sub categories database", fiber.StatusOK, ctx)
	}

	return response.InternalServiceDataResponse(productsSubCategories, fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) GetAllEventSubCategoriesByPagination(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.SubCategoryPaginationRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	var pagination models.Pagination
	pagination.Limit = request.Limit
	pagination.Sort = request.Sort
	pagination.Page = request.Page

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 04. GET EVENT SUB CATEGORIES AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	productSubCategoriesData, err := service.EventHubCategoriesSubCategoriesService.GetAllEventSubCategoriesByPagination(pagination, request.EventCategoryID, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(productSubCategoriesData, fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) UpdateEventSubCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventSubCategoriesUpdateRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	productSubCategoryID, conversionError := strconv.Atoi(ctx.FormValue("id"))
	if conversionError == nil {
		request.Id = uint64(productSubCategoryID)
	}
	request.EventSubCategoryName = ctx.FormValue("event_sub_category_name")
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.ValidateForUpdate(request, request.Id)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*------------------------------------------------------------
	  04. GET UPLOADED FILE
	 --------------------------------------------------------------*/
	file, fileErr := ctx.FormFile("image")
	savedPath := ""
	if file != nil {

		if fileErr != nil {
			return response.ErrorResponse(fileErr.Error(), fiber.StatusBadRequest, ctx)
		}

		/*------------------------------------------------------------
		  05. VALIDATING UPLOADED FILE TYPE
		 --------------------------------------------------------------*/
		allowedExtensions := []string{"png", "jpg", "jpeg"}
		uploadedFileExtension := strings.ToLower(strings.Split(file.Filename, ".")[1])
		if uploadedFileExtension != allowedExtensions[0] && uploadedFileExtension != allowedExtensions[1] &&
			uploadedFileExtension != allowedExtensions[2] {
			return response.ErrorResponse("Invalid file format. Supported. 'png', 'jpg', 'jpeg'", fiber.StatusBadRequest, ctx)
		}

		/*-----------------------------------------------
		  06. SAVING FILE TO THE LOCAL STORAGE LOCATION
		 ------------------------------------------------*/
		os.MkdirAll("public/products_subcategories/images/", os.ModePerm)
		savedPath = "/products_subcategories/images/" + strings.ToLower(random_string_generators.RandomString(20)) + "." + uploadedFileExtension
		pathData := "public" + savedPath

		saveErr := ctx.SaveFile(file, pathData)

		if saveErr != nil {
			return response.ErrorResponse(saveErr.Error(), fiber.StatusConflict, ctx)
		}

	}
	/*------------------------------------------------------------
	  07. ADD SAVED PATH TO THE REQUEST OBJECT AS URL
	 --------------------------------------------------------------*/
	if savedPath != "" {
		request.IconUrl = savedPath
		request.ImageStorage = "LOCAL"
	} else {
		request.IconUrl = ctx.FormValue("image")
		request.ImageStorage = "REMOTE"
	}
	/*-----------------------------------------------------------------
	 04. UPDATE EVENT SUB CATEGORY NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err := service.EventHubCategoriesSubCategoriesService.UpdateEventSubCategory(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event sub category name updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubCategoriesSubCategoriesController) DeleteEventSubCategory(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST
	---------------------------------------------------------*/
	var request categories_sub_categories.EventHubEventSubCategoriesGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse("Bad request", fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*-----------------------------------------------------------------
	 05. DELETE EVENT SUB CATEGORY AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubCategoriesSubCategoriesRepository.DeleteEventSubCategory(request.EventSubCategoryID)
	/*---------------------------------------------------------
	 06. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete product sub category", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event sub category deleted successfully!", fiber.StatusOK, ctx)
}
