package service

import (
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
)

var EventHubUsersManagementService = newEventHubUsersManagementService()

type eventHubUsersManagementService struct {
}

func newEventHubUsersManagementService() eventHubUsersManagementService {
	return eventHubUsersManagementService{}
}

func (_ eventHubUsersManagementService) RegisterUser(user models.EventHubUser) (*models.EventHubUser, error) {
	/*---------------------------------------------------------
	 01. CREATE USER AND GET DB RESPONSE AND CHECK AFFECTED ROWS
	----------------------------------------------------------*/
	createdUser, dbResponse := repositories.EventHubUsersManagementRepository.RegisterUser(&user)
	if dbResponse.RowsAffected == 0 {
		return nil, errors.New("Failed to register user!")
	}

	/*---------------------------------------------------------
	 02. ADD USER IN USER TOKEN TABLE
	----------------------------------------------------------*/
	return createdUser, EventHubUserTokenService.CreateUserTokenInDB(&user)
}

func (_ eventHubUsersManagementService) GetUserByPhone(phone string) (*models.EventHubUser, error) {
	user, dbResponse := repositories.EventHubUsersManagementRepository.FindOneByPhone(phone)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return user, errors.New("User with " + phone + " not found! ")
	}
	return user, nil
}

func (_ eventHubUsersManagementService) GetUsers(role string) ([]models.KataTiketiUserDTO, error) {
	car, dbResponse := repositories.EventHubUsersManagementRepository.GetUsers(role)
	if dbResponse.RowsAffected == 0 {
		// RETURN RESPONSE IF NO ROWS RETURNED
		return nil, errors.New("No Users found!")
	}
	return car, nil
}

func (_ eventHubUsersManagementService) ChangePassword(user models.EventHubUser, userID uint64) (string, error) {
	/*---------------------------------------------------------
	 01. INITIATING VALUES HOLDING MODEL VALUES BEFORE THE
	     UPDATES
	----------------------------------------------------------*/
	userFromDatabase, usDB := repositories.EventHubUsersManagementRepository.FindOne(userID)
	if usDB.RowsAffected == 0 {
		return "", errors.New("Internal server error")
	}
	userFromDatabase.Password = user.Password
	/*---------------------------------------------------------
	 02. UPDATING USER DETAILS TO THE DATABASE
	----------------------------------------------------------*/
	_, usrDB := repositories.EventHubUsersManagementRepository.UpdateWithID(userID, userFromDatabase)
	if usrDB.RowsAffected == 0 {
		return "", errors.New("Internal server error")
	}

	return "Password updated successful", nil
}
