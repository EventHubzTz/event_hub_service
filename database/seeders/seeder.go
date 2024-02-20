package seeders

import (
	"github.com/EventHubzTz/event_hub_service/database"
	"github.com/EventHubzTz/event_hub_service/utils"
)

func Seed() {
	db := database.DB()
	/*---------------------------------------------------
	  01. REGISTERING SEEDER FOR MICROSERVICE REQUEST ID
	 ----------------------------------------------------*/
	err := EventHubRequestIDTableSeeder(db)
	if err != nil {
		utils.ErrorPrint(err.Error())
		// return
	}
	utils.SuccessPrint("Seed Request ID Successfully!")

	/*---------------------------------------------------
	  02. REGISTERING SEEDER FOR EXTERNAL OPERATION SETUP
	 ----------------------------------------------------*/
	err = EventHubExternalOperationSetupTableSeeder(db)
	if err != nil {
		utils.ErrorPrint(err.Error())
		//  return
	}
	utils.SuccessPrint("Seed External Setup Successfully!")
}
