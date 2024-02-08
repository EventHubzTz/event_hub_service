package repositories

import (
	"github.com/EventHubzTz/event_hub_service/database"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() {
	db = database.DB()
}
