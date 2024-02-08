package middlewares

import (
	"fmt"
	"os"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ApiAuth(ctx *fiber.Ctx) error {
	/*-----------------------------------------------
	 01. GET VALUE OF THE 'Authorization' HEADER
	     SUBMITTED IN INCOMING REQUEST
	     (BEARER TOKEN)
	------------------------------------------------*/
	headerBearToken := ctx.Get("event-hub-token-auth")
	/*-----------------------------------------------
	     02. EXTRACTING TOKEN VALUE FROM THE BEARER TOKEN
		     VALUE (SPLITTING SPACE)
		-------------------------------------------------*/
	// bearArray := strings.Split(headerBearToken, " ")
	// if len(bearArray) != 2 {
	// 	return response.ErrorResponseStr("Unexpected error", fiber.StatusBadRequest, ctx)
	// }

	// bearToken := bearArray[1]
	fmt.Println(">>>>>>>>>>>>>>>>>" + headerBearToken)
	/*-------------------------------------------------
	03. JWT CLAIM DETAILS BASED ON VALID ACCESS TOKEN
	--------------------------------------------------*/
	token, err := jwt.ParseWithClaims(headerBearToken, &models.EventHubUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("APP_KEY")), nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return response.ErrorResponse("error Unauthorized Access", fiber.StatusUnauthorized, ctx)
	}
	/*-------------------------------------------------
	04. VERIFICATION OF JWT CLAIM EXTRACTED BY THE
	    TOKEN
	--------------------------------------------------*/
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return response.ErrorResponse("Unauthorized Access", fiber.StatusUnauthorized, ctx)
	}
	/*-------------------------------------------------
	 05. EXTRACTION OF THE USER CLAIM FROM THE JWT
	     CLAIM
	--------------------------------------------------*/
	claims := token.Claims.(*models.EventHubUserClaims)
	ctx.Locals("user", claims)
	return ctx.Next()
}
