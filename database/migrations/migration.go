package migrations

import (
	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/database"
	"github.com/EventHubzTz/event_hub_service/utils"
)

func Migrate() {
	db := database.DB()
	err := db.AutoMigrate(models.Tables...)
	if err != nil {
		utils.ErrorPrint(err.Error())
		return
	}
	utils.SuccessPrint("Migrated Successfully!")
}

func DropTables() {
	var db = database.DB()
	err := db.Migrator().DropTable(models.Tables...)
	if err != nil {
		utils.ErrorPrint(err.Error())
		return
	}
	utils.SuccessPrint("Dropping Tables Successfully!")
}
