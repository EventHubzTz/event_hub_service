package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/EventHubzTz/event_hub_service/app/http/requests/events"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/EventHubzTz/event_hub_service/service"
	"github.com/EventHubzTz/event_hub_service/utils/date_utils"
	"github.com/EventHubzTz/event_hub_service/utils/process_image"
	"github.com/EventHubzTz/event_hub_service/utils/random_string_generators"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/EventHubzTz/event_hub_service/utils/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var EventHubEventsManagementController = newEventHubEventsManagementController()

type eventHubEventsManagementController struct {
}

func newEventHubEventsManagementController() eventHubEventsManagementController {
	return eventHubEventsManagementController{}
}

func (c eventHubEventsManagementController) AddEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	userFromLocal := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	request.UserID = userFromLocal.Id
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*--------------------------------------------------------------------
	 04. ADD EVENT
	-----------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.AddEvent(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Event added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) AddEventImage(ctx *fiber.Ctx) error {
	/*--------------------------------------------------------------------
	01. INITIALIZING VARIABLE FOR REQUEST OF STORING EVENT COVER IMAGE
	----------------------------------------------------------------------*/
	var request events.EventHubEventImageRequest

	/*-----------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING MULTIPART/FORM REQUEST
	-------------------------------------------------------------*/
	productID, errContentID := strconv.Atoi(ctx.FormValue("event_id"))
	if errContentID == nil {
		request.EventID = uint64(productID)
	}
	request.FileType = "IMAGE"
	request.ImageStorage = "LOCAL"
	/*---------------------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS IN A REQUEST
	-----------------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	/*----------------------------------------------------------
	 04. VALIDATING EVENT IMAGE FILE EXTENSION IF SUPPORTED
	------------------------------------------------------------*/
	file, er := ctx.FormFile("image")
	if er != nil {
		return response.ErrorResponse(er.Error(), fiber.StatusBadRequest, ctx)
	}
	allowedExtensions := []string{"png", "jpg", "jpeg", "mp4", "mkv", "svg", "webp"}
	uploadedFileExtension := strings.ToLower(strings.Split(file.Filename, ".")[len(strings.Split(file.Filename, "."))-1])
	if uploadedFileExtension != allowedExtensions[0] &&
		uploadedFileExtension != allowedExtensions[1] &&
		uploadedFileExtension != allowedExtensions[2] &&
		uploadedFileExtension != allowedExtensions[3] &&
		uploadedFileExtension != allowedExtensions[4] &&
		uploadedFileExtension != allowedExtensions[5] &&
		uploadedFileExtension != allowedExtensions[6] {
		return response.ErrorResponse("Invalid file format. Supported. 'png', 'jpg', 'jpeg' 'mp4', 'mkv', 'webp'", fiber.StatusBadRequest, ctx)
	}

	if file.Size > (50 * 1024 * 1024) {
		return response.ErrorResponse("The file is too large. Expected: "+strconv.Itoa(int(50))+"MB provided: "+strconv.Itoa(int(file.Size/(1024*1024)))+"MB", fiber.StatusBadRequest, ctx)
	}

	/*----------------------------------------
	 05. CHECK IF EVENT EXIST IN THE SYSTEM
	------------------------------------------*/
	_, eventError := service.EventHubEventsManagementService.GetEvent(request.EventID)
	if eventError != nil {
		return response.ErrorResponse(eventError.Error(), fiber.StatusBadRequest, ctx)
	}

	/*---------------------------------------------------
	 06.  CHECKING IF IMAGE COVER REACH MAX LIMIT (5)
	----------------------------------------------------*/
	err := service.EventHubEventsManagementService.CheckIfEventReachMaxCoverImageLimit(request.EventID)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusConflict, ctx)
	}
	/*---------------------------------------------------
	 07. SAVING IMAGE FILE TO THE LOCAL STORAGE LOCATION
	----------------------------------------------------*/
	os.MkdirAll("public/products/images/", os.ModePerm)
	id := uuid.New()
	uploadPath := "/products/images/" + strings.ToLower(random_string_generators.RandomString(20)) + id.String() + "." + uploadedFileExtension
	pathData := "public" + uploadPath
	err = ctx.SaveFile(file, pathData)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusConflict, ctx)
	}

	/*---------------------------------------------------
	 08. CALCULATE IMAGE ASPECT RATIO
	----------------------------------------------------*/
	aspectRatio, _ := process_image.ProcessImage.GetImageAspectRatios(pathData)
	request.AspectRatios = aspectRatio
	request.ImagePath = uploadPath

	if uploadedFileExtension == "mp4" {
		fileName := strings.ToLower(random_string_generators.RandomString(20)) + id.String()
		os.MkdirAll("public/products/streamable_videos/"+fileName+"/", os.ModePerm)
		os.MkdirAll("public/products/thumbnails/", os.ModePerm)
		outputHLS := "/products/streamable_videos/" + fileName + "/" + fileName + ".m3u8"
		outputThumbnail := "/products/thumbnails/" + fileName + ".jpg"
		request.FileType = "VIDEO"
		request.ImagePath = outputHLS
		request.VideoUrl = uploadPath
		request.ThumbunailUrl = outputThumbnail

		go func() {
			//Generate Thumbnail
			publicFolderPath := "/go/event_hub_service/public"
			thumbnailPath := publicFolderPath + outputThumbnail
			thumbnailTime := "00:00:05" // Time at which to capture the thumbnail (e.g., 5 seconds)
			err = process_image.GenerateThumbnail("/go/event_hub_service/"+pathData, thumbnailPath, thumbnailTime)

			// Convert the uploaded video to HLS
			cmd := exec.Command("ffmpeg", "-i", "/go/event_hub_service/"+pathData, "-c:v", "libx264", "-c:a", "aac", "-f", "hls", "-hls_time", "10", "-hls_list_size", "0", publicFolderPath+outputHLS)
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Print(err)
			}
		}()
	}

	/*---------------------------------------------
	 09. ADD EVENT IMAGE AND GET ERROR IF IS AVAILABLE
	-----------------------------------------------*/
	err = service.EventHubEventsManagementService.AddEventImage(request.ToModel())

	/*---------------------------------------------------------
	 10. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		process_image.ProcessImage.DeleteImage(pathData)
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	/*---------------------------------------------------------
	 11. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.SuccessResponse("Event image added successfully on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetEvents(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
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
	 04. GET EVENTS AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	user := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	events, err := service.EventHubEventsManagementService.GetEvents(pagination, user.Role, request.Query, user.Id, request.ProductCategoryID, request.ProductSubCategoryID)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(events, fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetEvent(ctx *fiber.Ctx) error {
	/*--------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST UPDATING PASSWORD
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*---------------------------------------------------------
	 04. GET THE EVENT FROM THE DATABASE USING EVENT ID
	----------------------------------------------------------*/
	event, eventError := service.EventHubEventsManagementService.GetEvent(request.EventID)
	if eventError != nil {
		return response.ErrorResponse(eventError.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 09. IF ALL THIS WENT WELL THEN RETURN SUCCESS
	----------------------------------------------------------*/
	return response.InternalServiceDataResponse(event, fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) UpdateEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubUpdateEventRequest
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
	 04. UPDATE EVENT NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.UpdateEvent(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) DeleteEventImage(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
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
	 04. DELETE EVENT IMAGE AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubEventsManagementRepository.DeleteEventImage(request.EventID)
	/*---------------------------------------------------------
	 05. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete product image", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event image deleted successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) DeleteEvent(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
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
	 04. DELETE EVENT AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubEventsManagementRepository.DeleteEvent(request.EventID)
	/*---------------------------------------------------------
	 05. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete event", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event deleted successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) AddEventPackage(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventPackageRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*----------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	------------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	/*--------------------------------------------------------------------
	 04. ADD EVENT PACKAGE
	-----------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.AddEventPackage(request.ToModel())
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusInternalServerError, ctx)
	}

	return response.SuccessResponse("Event package added successful on "+date_utils.GetNowString(), fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetAllEventPackages(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventPackagesGetsRequest
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
	 04. GET EVENT PACKAGES AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	events, err := service.EventHubEventsManagementService.GetAllEventPackages(pagination, request.EventID, request.Query)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusNotFound, ctx)
	}
	return response.InternalServiceDataResponse(events, fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) UpdateEventPackage(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubUpdateEventPackageRequest
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
	 04. UPDATE EVENT NAME AND GET ERROR IF IS AVAILABLE
	-------------------------------------------------------------------*/
	err = service.EventHubEventsManagementService.UpdateEventPackage(request.ToModel(), request.Id)
	/*---------------------------------------------------------
	 05. CHECK IF ERROR IS AVAILABLE AND RETURN ERROR RESPONSE
	----------------------------------------------------------*/
	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event updated successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) DeleteEventPackage(ctx *fiber.Ctx) error {
	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF GETTING
	     CONTENTS
	---------------------------------------------------------*/
	var request events.EventHubEventGetRequest
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
	 04. DELETE EVENT PACKAGE AND GET RESPONSE IF IS AVAILABLE
	-------------------------------------------------------------------*/
	dbResponse := repositories.EventHubEventsManagementRepository.DeleteEventPackage(request.EventID)
	/*---------------------------------------------------------
	 05. CHECK IF ROW IS AFFECTED AND RETURN RESPONSE
	----------------------------------------------------------*/
	if dbResponse.RowsAffected == 0 {
		return response.ErrorResponse("Failed to delete event package", fiber.StatusBadRequest, ctx)
	}

	return response.SuccessResponse("Event package deleted successfully!", fiber.StatusOK, ctx)
}

func (c eventHubEventsManagementController) GetDashboardStatistics(ctx *fiber.Ctx) error {
	type OrdersStatisticsForVendorRequest struct {
		From string `json:"from" validate:"required"`
		To   string `json:"to" validate:"required"`
	}

	/*-------------------------------------------------------
	 01. INITIATING VARIABLE FOR THE REQUEST OF REGISTERING
	     USER
	---------------------------------------------------------*/
	var request OrdersStatisticsForVendorRequest
	/*---------------------------------------------------------
	 02. PARSING THE BODY OF THE INCOMING REQUEST
	----------------------------------------------------------*/
	err := ctx.BodyParser(&request)

	if err != nil {
		return response.ErrorResponse(err.Error(), fiber.StatusBadRequest, ctx)
	}
	/*---------------------------------------------------------
	 03. VALIDATING THE INPUT FIELDS OF THE PASSED PARAMETERS
	     IN A REQUEST
	----------------------------------------------------------*/
	errors := validation.Validate(request)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user := service.EventHubUserTokenService.GetUserFromLocal(ctx)
	return response.MapDataResponse(service.EventHubEventsManagementService.GetDashboardStatistics(user.Role, request.From, request.To, user.Id), fiber.StatusOK, ctx)
}
