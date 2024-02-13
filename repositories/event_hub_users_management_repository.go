package repositories

import (
	"database/sql"
	"errors"

	"github.com/EventHubzTz/event_hub_service/app/helpers"
	"github.com/EventHubzTz/event_hub_service/app/models"
	"gorm.io/gorm"
)

var EventHubUsersManagementRepository = newEventHubUsersManagementRepository()

type eventHubUsersManagementRepository struct {
}

func newEventHubUsersManagementRepository() eventHubUsersManagementRepository {
	return eventHubUsersManagementRepository{}
}

func (_ eventHubUsersManagementRepository) RegisterUser(user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Create(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) FindOneByPhone(phone string) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Model(models.EventHubUser{}).Where("phone_number = ?", phone).Find(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) FindUserUsingPhoneNumber(phoneNumber string) (*models.EventHubUserDTO, *gorm.DB) {
	var user *models.EventHubUserDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QuerySpecificUserDetailsUsingPhoneNumber(), phoneNumber).Find(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) GetUsers(role string) ([]models.EventHubUserDTO, *gorm.DB) {
	var renting []models.EventHubUserDTO
	sRDB := db.Raw(helpers.EventHubQueryBuilder.QueryGetUsers(role), role).Find(&renting)
	return renting, sRDB
}

func (_ eventHubUsersManagementRepository) FindUserById(userId uint64) *models.EventHubUser {
	var user *models.EventHubUser
	dbErr := db.Where("id = ?", userId).Find(&user)
	if dbErr.RowsAffected == 0 {
		return nil
	}
	return user
}

func (_ eventHubUsersManagementRepository) FindOne(id uint64) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryUserDetails(), id).Find(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) UpdateWithID(id uint64, user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Model(models.EventHubUser{}).Where("id = ?", id).Updates(&user)
	return user, urDB
}

func (_ eventHubUsersManagementRepository) SaveUserOTPCode(otpCode *models.EventHubUserOTPCode) (*models.EventHubUserOTPCode, *gorm.DB) {
	urDB := db.Create(&otpCode)
	return otpCode, urDB
}

func (_ eventHubUsersManagementRepository) SaveUserOTPCodeMessage(otpCodeMessage *models.EventHubOTPCodeMessage) (*models.EventHubOTPCodeMessage, *gorm.DB) {
	urDB := db.Create(&otpCodeMessage)
	return otpCodeMessage, urDB
}

func (_ eventHubUsersManagementRepository) SaveUserOTPCodeMessageResponse(response *models.EventHubOTPMessageResponse) (*models.EventHubOTPMessageResponse, *gorm.DB) {
	urDB := db.Create(&response)
	return response, urDB
}

func (_ eventHubUsersManagementRepository) VerifyPhoneNumberAndOTPCode(otpCode models.EventHubUserOTPCode) bool {
	var otpCodeDetails *models.EventHubUserOTPCode
	urDB := db.Where("phone = @phone AND otp = @otp", sql.Named("phone", otpCode.Phone),
		sql.Named("otp", otpCode.OTP)).Find(&otpCodeDetails)
	if urDB.RowsAffected == 0 {
		return false
	}
	return true
}

func (_ eventHubUsersManagementRepository) VerifyUserPhoneNumber(phoneNumber string, userID uint64) bool {
	var user *models.EventHubUser
	urDB := db.Where("phone_number = @phone_number AND id = @id", sql.Named("phone_number", phoneNumber),
		sql.Named("id", userID)).Find(&user)
	if urDB.RowsAffected == 0 {
		return false
	}
	return true
}

func (_ eventHubUsersManagementRepository) VerifyUserPhoneNumberInvalidStatus(phoneNumber string, userID uint64) bool {
	var user *models.EventHubUser
	urDB := db.Where("phone_number = @phone_number AND id = @id AND is_valid_number = @is_valid_number", sql.Named("phone_number", phoneNumber),
		sql.Named("id", userID), sql.Named("is_valid_number", 0)).Find(&user)
	if urDB.RowsAffected == 0 {
		return false
	}
	return true
}

func (_ eventHubUsersManagementRepository) UpdateUserPhoneNumberValidStatus(user *models.EventHubUser) error {
	urDB := db.Model(&user).Update("is_valid_number", 1)
	if urDB.RowsAffected == 0 {
		return errors.New("Failed to verify mobile phone otp code")
	}
	return nil
}

func (_ eventHubUsersManagementRepository) FindForgetPasswordOTPDetails(userID uint64) *models.AFYAAPPForgotPasswordOTP {
	var userForgotPasswordDetails *models.AFYAAPPForgotPasswordOTP
	urDB := db.Where("user_id = ?", userID).Find(&userForgotPasswordDetails)
	if urDB.RowsAffected == 0 {
		return nil
	}
	return userForgotPasswordDetails
}

func (_ eventHubUsersManagementRepository) CreateOTPCodeForForgotPassword(userForgotPasswordOTP *models.AFYAAPPForgotPasswordOTP) (*models.AFYAAPPForgotPasswordOTP, *gorm.DB) {
	urDB := db.Create(&userForgotPasswordOTP)
	return userForgotPasswordOTP, urDB
}

func (_ eventHubUsersManagementRepository) UpdateOTPCodeForForgotPassword(userForgotPasswordOTP *models.AFYAAPPForgotPasswordOTP) (*models.AFYAAPPForgotPasswordOTP, *gorm.DB) {
	urDB := db.Save(&userForgotPasswordOTP)
	return userForgotPasswordOTP, urDB
}
