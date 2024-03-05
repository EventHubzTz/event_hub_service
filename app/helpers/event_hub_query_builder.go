package helpers

import (
	"math"
	"os"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/database"
	"github.com/EventHubzTz/event_hub_service/utils/constants"
	"gorm.io/gorm"
)

var EventHubQueryBuilder = newEventHubQueryBuilder()

type eventHubQueryBuilder struct {
}

func newEventHubQueryBuilder() eventHubQueryBuilder {
	return eventHubQueryBuilder{}
}

func paginate(value interface{}, pagination *models.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var totalRows int64

		// Create a new session for counting the affected rows
		countDB := db.Session(&gorm.Session{})

		// Apply the offset, limit, and order conditions to the query
		countDB = countDB.Model(value)

		// Count the affected rows
		countDB.Count(&totalRows)

		pagination.TotalResults = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
		pagination.TotalPages = totalPages

		// Return the query with offset, limit, and order conditions
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func (q eventHubQueryBuilder) QueryMicroServiceRequestIDActiveKey() string {
	return "SELECT t1.id,t1.request_id FROM event_hub_request_ids t1 " +
		"ORDER BY t1.id DESC LIMIT 1"
}

func (q eventHubQueryBuilder) QueryConfigurations() string {
	return "SELECT t1.*," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at FROM event_hub_configurations t1 "
}

func (q eventHubQueryBuilder) QueryGetUsers(pagination models.Pagination, role, query string) (models.Pagination, *gorm.DB) {
	baseUrl := os.Getenv("APP_URL")

	var events []models.EventHubUserDTO

	clDB := database.DB().Scopes(paginate([]models.EventHubUser{}, &pagination, database.DB())).
		Table("event_hub_users as t1").
		Select(
			"t1.*",
			"CASE t1.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',t1.profile_image) ELSE t1.profile_image END as profile_image,"+
				"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		)
	if role != "" {
		clDB = clDB.Where("t1.role = ?", role)
	}
	if query != "%%" {
		clDB = clDB.Where(
			"concat(t1.first_name,' ',t1.last_name,' ',t1.email,' ',t1.phone_number,' ',t1.gender,' ',DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p')) like ? ",
			query,
		)
	}
	clDB = clDB.Find(&events)
	pagination.Results = events
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryUserDetails() string {
	baseUrl := os.Getenv("APP_URL")

	return "SELECT t1.*," +
		"CASE t1.image_storage WHEN 'LOCAL' THEN CONCAT('" + baseUrl + "',t1.profile_image) ELSE t1.profile_image END as profile_image," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at " +
		"FROM event_hub_users t1 " +
		"WHERE t1.active = true and t1.id = ?"
}

func (q eventHubQueryBuilder) QuerySpecificUserDetailsUsingPhoneNumber() string {
	baseUrl := os.Getenv("APP_URL")
	return "SELECT DISTINCT t1.*," +
		"CASE t1.image_storage WHEN 'LOCAL' THEN CONCAT('" + baseUrl + "',t1.profile_image) ELSE t1.profile_image END as profile_image," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at " +
		"FROM event_hub_users t1 " +
		"WHERE t1.phone_number =  ? "
}

func (q eventHubQueryBuilder) QueryAllEventCategories(pagination models.Pagination, query string) (models.Pagination, *gorm.DB) {
	var eventCategories []models.EventHubEventCategoriesDTO
	baseUrl := os.Getenv("APP_URL")

	clDB := database.DB().Scopes(paginate([]models.EventHubEventCategories{}, &pagination, database.DB())).
		Table("event_hub_event_categories as t1").
		Select(
			"t1.*",
			"CONCAT(CASE t1.image_storage WHEN 'LOCAL' THEN '"+baseUrl+"' ELSE '' END, t1.icon_url) as icon_url",
			"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		)
	if query != "%%" {
		clDB = clDB.Where("t1.id LIKE ? OR "+
			"t1.event_category_name LIKE ? ",
			query,
			query,
		)
	}
	clDB = clDB.Find(&eventCategories)
	pagination.Results = eventCategories
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryAllEventSubCategories(pagination models.Pagination, eventCategoryId uint64, query string) (models.Pagination, *gorm.DB) {
	var eventSubCategories []models.EventHubEventSubCategoriesDTO
	baseUrl := os.Getenv("APP_URL")

	clDB := database.DB().Scopes(paginate([]models.EventHubEventSubCategories{}, &pagination, database.DB())).
		Table("event_hub_event_subcategories as t1").
		Select(
			"t1.*",
			"CONCAT(CASE t1.image_storage WHEN 'LOCAL' THEN '"+baseUrl+"' ELSE '' END, t1.icon_url) as icon_url",
			"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		).Where("t1.event_category_id = ?", eventCategoryId)
	if query != "%%" {
		clDB = clDB.Where("t1.id LIKE ? OR "+
			"t1.event_sub_category_name LIKE ? ",
			query,
			query,
		)
	}
	clDB = clDB.Find(&eventSubCategories)
	pagination.Results = eventSubCategories
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryGetEvents(pagination models.Pagination, role, query string, userID, eventCategoryId, eventSubCategoryId uint64) (models.Pagination, *gorm.DB) {
	var events []models.EventHubEventDTO
	baseUrl := os.Getenv("APP_URL")

	clDB := database.DB().Scopes(paginate([]models.EventHubEvent{}, &pagination, database.DB())).
		Table("event_hub_events as t1").
		Joins("LEFT JOIN event_hub_event_categories t2 on t1.event_category_id = t2.id").
		Joins("LEFT JOIN event_hub_event_subcategories t3 on t1.event_sub_category_id = t3.id").
		Joins("LEFT JOIN event_hub_users t4 on t1.user_id = t4.id").
		Select(
			"t1.*",
			"t2.event_category_name",
			"t3.event_sub_category_name",
			"CONCAT(t4.first_name, ' ', t4.last_name) as event_owner",
			"CASE t4.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',t4.profile_image) ELSE t4.profile_image END as event_owner_profile,"+
				"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		).
		Preload("EventFiles", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_hub_event_images").
				Select(
					"event_hub_event_images.*",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.image_url) ELSE event_hub_event_images.image_url END image_url",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.video_url) ELSE event_hub_event_images.video_url END video_url",
					"CASE event_hub_event_images.image_storage WHEN 'LOCAL' THEN CONCAT('"+baseUrl+"',event_hub_event_images.thumbunail_url) ELSE event_hub_event_images.thumbunail_url END thumbunail_url",
				)
		}).
		Preload("EventPackages", func(db *gorm.DB) *gorm.DB {
			return db.Table("event_hub_event_packages").
				Select(
					"event_hub_event_packages.*",
				)
		})
	if role == constants.EventPlanner {
		clDB = clDB.Where("t1.user_id = ?", userID)
	}
	if eventCategoryId != 0 {
		clDB = clDB.Where("t1.event_category_id = ?", eventCategoryId)
	}
	if eventSubCategoryId != 0 {
		clDB = clDB.Where("t1.event_sub_category_id = ?", eventSubCategoryId)
	}
	if query != "%%" {
		clDB = clDB.Where(
			"concat(t1.event_name,' ',t1.event_location,' ',t1.event_description,' ',DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p')) like ? ",
			query,
		)
	}
	clDB = clDB.Find(&events)
	pagination.Results = events
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryAllEventPackages(pagination models.Pagination, eventID uint64, query string) (models.Pagination, *gorm.DB) {
	var eventPackages []models.EventHubEventPackagesDTO

	clDB := database.DB().Scopes(paginate([]models.EventHubEventPackages{}, &pagination, database.DB())).
		Table("event_hub_event_packages as t1").
		Select(
			"t1.*",
			"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at",
			"DATE_FORMAT(t1.updated_at, '%W, %D %M %Y %h:%i:%S%p') as updated_at",
		).
		Where("t1.event_id = ?", eventID)
	if query != "%%" {
		clDB = clDB.Where("t1.id LIKE ? OR "+
			"t1.package_name LIKE ? OR "+
			"t1.amount LIKE ? ",
			query,
			query,
			query,
		)
	}
	clDB = clDB.Find(&eventPackages)
	pagination.Results = eventPackages
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryEventDetails() string {

	return "SELECT t1.*," +
		"DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p') as created_at " +
		"FROM event_hub_events t1 " +
		"WHERE t1.id = ?"
}

func (q eventHubQueryBuilder) QueryPaymentTransactions(pagination models.Pagination, role, query, status, phoneNumber string, userID uint64) (models.Pagination, *gorm.DB) {
	var events []models.EventHubPaymentTransactionsDTO

	clDB := database.DB().Scopes(paginate([]models.EventHubPaymentTransactions{}, &pagination, database.DB())).
		Table("event_hub_payment_transactions as t1").
		Joins("LEFT JOIN event_hub_users t2 on t1.user_id = t2.id").
		Joins("LEFT JOIN event_hub_events t3 on t1.event_id = t3.id").
		Select(
			"t1.*",
			"CONCAT(t2.first_name, ' ', t2.last_name) as full_name",
			"t3.event_name",
			"CONCAT(YEAR(CURDATE()) - YEAR(STR_TO_DATE(t1.date_of_birth,'%d/%m/%Y')), ' Year(s)') as age",
		)
	if role == constants.EventPlanner {
		clDB = clDB.Where("t3.user_id = ?", userID)
	}
	if role == constants.NormalUser {
		clDB = clDB.Where("t1.user_id = ?", userID)
	}
	if status != "" {
		clDB = clDB.Where("t1.payment_status = ?", status)
	}
	if phoneNumber != "" {
		clDB = clDB.Where("t1.phone_number = ?", phoneNumber).
			Where("t1.payment_status = ?", constants.Completed)
	}
	if query != "%%" {
		clDB = clDB.Where(
			"concat(t1.order_id,' ',t1.transaction_id,' ',t1.phone_number,' ',t1.amount,' ',t1.currency,' ',t1.provider,' ',t1.payment_status,' ',t2.first_name,' ',t2.last_name,' ',t3.event_name,' ',DATE_FORMAT(t1.created_at, '%W, %D %M %Y %h:%i:%S%p')) like ? ",
			query,
		)
	}
	clDB = clDB.Find(&events)
	pagination.Results = events
	return pagination, clDB

}

func (q eventHubQueryBuilder) QueryGetDashboardStatistics() string {
	return "SELECT " +
		"(SELECT COUNT(*) FROM event_hub_users) AS total_users," +
		"(SELECT COUNT(*) FROM event_hub_events) AS total_events," +
		"(SELECT SUM(amount) FROM event_hub_payment_transactions WHERE payment_status='COMPLETED') AS total_amount," +
		"(SELECT SUM(amount * 0.3) FROM event_hub_payment_transactions WHERE payment_status='COMPLETED') AS agregator_collection," +
		"(SELECT SUM(amount * 0.3) FROM event_hub_payment_transactions WHERE payment_status='COMPLETED') AS system_collection," +
		"(SELECT SUM(amount * 0.94) FROM event_hub_payment_transactions WHERE payment_status='COMPLETED') AS remained_collection"
}

func (q eventHubQueryBuilder) QueryGetDashboardStatisticsForEventPlanner() string {
	return "SELECT " +
		"(SELECT COUNT(*) FROM event_hub_events WHERE user_id = ?) AS total_events," +
		"(SELECT SUM(amount) FROM event_hub_payment_transactions LEFT JOIN event_hub_events ON event_hub_payment_transactions.event_id = event_hub_events.id WHERE event_hub_payment_transactions.payment_status='COMPLETED' AND event_hub_events.user_id = ?) AS total_amount"
}
