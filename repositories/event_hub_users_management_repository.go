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

func (r eventHubUsersManagementRepository) RegisterUser(user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Create(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) FindOneByEmailPhone(emailPhone string) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Model(models.EventHubUser{}).Where("phone_number = ?", emailPhone).Or("email = ?", emailPhone).Find(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) FindUserUsingPhoneNumber(phoneNumber string) (*models.EventHubUserDTO, *gorm.DB) {
	var user *models.EventHubUserDTO
	urDB := db.Raw(helpers.EventHubQueryBuilder.QuerySpecificUserDetailsUsingPhoneNumber(), phoneNumber).Find(&user)
	return user, urDB
}

func (r_ eventHubUsersManagementRepository) GetUsers(pagination models.Pagination, role, query string) (models.Pagination, *gorm.DB) {

	users, urDB := helpers.EventHubQueryBuilder.QueryGetUsers(pagination, role, query)

	return users, urDB
}

func (r eventHubUsersManagementRepository) FindUserById(userId uint64) *models.EventHubUser {
	var user *models.EventHubUser
	dbErr := db.Where("id = ?", userId).Find(&user)
	if dbErr.RowsAffected == 0 {
		return nil
	}
	return user
}

func (r eventHubUsersManagementRepository) FindOne(id uint64) (*models.EventHubUser, *gorm.DB) {
	var user *models.EventHubUser
	urDB := db.Raw(helpers.EventHubQueryBuilder.QueryUserDetails(), id).Find(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) UpdateWithID(id uint64, user *models.EventHubUser) (*models.EventHubUser, *gorm.DB) {
	urDB := db.Model(models.EventHubUser{}).Where("id = ?", id).Updates(&user)
	return user, urDB
}

func (r eventHubUsersManagementRepository) SaveUserOTPCode(otpCode *models.EventHubUserOTPCode) (*models.EventHubUserOTPCode, *gorm.DB) {
	urDB := db.Create(&otpCode)
	return otpCode, urDB
}

func (r eventHubUsersManagementRepository) SaveUserOTPCodeMessage(otpCodeMessage *models.EventHubOTPCodeMessage) (*models.EventHubOTPCodeMessage, *gorm.DB) {
	urDB := db.Create(&otpCodeMessage)
	return otpCodeMessage, urDB
}

func (r eventHubUsersManagementRepository) SaveUserOTPCodeMessageResponse(response *models.EventHubOTPMessageResponse) (*models.EventHubOTPMessageResponse, *gorm.DB) {
	urDB := db.Create(&response)
	return response, urDB
}

func (r eventHubUsersManagementRepository) VerifyPhoneNumberAndOTPCode(otpCode models.EventHubUserOTPCode) bool {
	var otpCodeDetails *models.EventHubUserOTPCode
	urDB := db.Where("phone = @phone AND otp = @otp", sql.Named("phone", otpCode.Phone),
		sql.Named("otp", otpCode.OTP)).Find(&otpCodeDetails)
	return urDB.RowsAffected != 0
}

func (r eventHubUsersManagementRepository) VerifyUserPhoneNumber(phoneNumber string, userID uint64) bool {
	var user *models.EventHubUser
	urDB := db.Where("phone_number = @phone_number AND id = @id", sql.Named("phone_number", phoneNumber),
		sql.Named("id", userID)).Find(&user)
	return urDB.RowsAffected != 0
}

func (r eventHubUsersManagementRepository) VerifyUserPhoneNumberInvalidStatus(phoneNumber string, userID uint64) bool {
	var user *models.EventHubUser
	urDB := db.Where("phone_number = @phone_number AND id = @id AND is_valid_number = @is_valid_number", sql.Named("phone_number", phoneNumber),
		sql.Named("id", userID), sql.Named("is_valid_number", 0)).Find(&user)
	return urDB.RowsAffected != 0
}

func (r eventHubUsersManagementRepository) UpdateUserPhoneNumberValidStatus(user *models.EventHubUser) error {
	urDB := db.Model(&user).Update("is_valid_number", 1)
	if urDB.RowsAffected == 0 {
		return errors.New("failed to verify mobile phone otp code")
	}
	return nil
}

func (r eventHubUsersManagementRepository) FindForgetPasswordOTPDetails(userID uint64) *models.EventHubForgotPasswordOTP {
	var userForgotPasswordDetails *models.EventHubForgotPasswordOTP
	urDB := db.Where("user_id = ?", userID).Find(&userForgotPasswordDetails)
	if urDB.RowsAffected == 0 {
		return nil
	}
	return userForgotPasswordDetails
}

func (r eventHubUsersManagementRepository) CreateOTPCodeForForgotPassword(userForgotPasswordOTP *models.EventHubForgotPasswordOTP) (*models.EventHubForgotPasswordOTP, *gorm.DB) {
	urDB := db.Create(&userForgotPasswordOTP)
	return userForgotPasswordOTP, urDB
}

func (r eventHubUsersManagementRepository) UpdateOTPCodeForForgotPassword(userForgotPasswordOTP *models.EventHubForgotPasswordOTP) (*models.EventHubForgotPasswordOTP, *gorm.DB) {
	urDB := db.Save(&userForgotPasswordOTP)
	return userForgotPasswordOTP, urDB
}

func (r eventHubUsersManagementRepository) DeleteUser(userId uint64) *gorm.DB {
	sRDB := db.Where("id = ? ", userId).Delete(models.EventHubUser{})
	return sRDB
}
