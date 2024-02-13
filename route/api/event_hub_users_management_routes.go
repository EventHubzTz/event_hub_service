package api

import (
	"github.com/EventHubzTz/event_hub_service/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventHubUsersManagementRoutes(route fiber.Router) {
	route.Post("/register/user", controllers.EventHubUsersManagementController.RegisterUser)
	route.Post("/resend/otp", controllers.EventHubUsersManagementController.ResendOTPCode)
	route.Post("/verify/phone", controllers.EventHubUsersManagementController.VerifyPhoneNumberUsingOTP)
	route.Post("/login/user", controllers.EventHubUsersManagementController.LoginUser)
	route.Post("/get/users", controllers.EventHubUsersManagementController.GetUsers)
	route.Post("/change/password", controllers.EventHubUsersManagementController.ChangePassword)
	route.Post("/generate/forgot/password/otp", controllers.EventHubUsersManagementController.GenerateForgotPasswordOtp)
	route.Post("/verify/otp/reset/password", controllers.EventHubUsersManagementController.VerifyOTPResetPassword)
	route.Post("/update/password", controllers.EventHubUsersManagementController.UpdatePassword)
}

func AuthenticatedEventHubUsersManagementRoutes(route fiber.Router) {

}
