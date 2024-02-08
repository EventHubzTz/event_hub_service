package service

import (
	"errors"
	"os"
	"time"

	"github.com/EventHubzTz/event_hub_service/app/models"
	"github.com/EventHubzTz/event_hub_service/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var EventHubUserTokenService = newEventHubUserTokenService()

type eventHubUserTokenService struct {
}

func newEventHubUserTokenService() eventHubUserTokenService {
	return eventHubUserTokenService{}
}

func (uts eventHubUserTokenService) generateJwtToken(user *models.EventHubUser) (uint64, string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	var value uint64 = 31536000 //YEAR
	expireAt := time.Now().Add(time.Second * time.Duration(value)).Unix()
	jwtToken.Claims = &models.EventHubUserClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
		TokenType:    "user-token",
		EventHubUser: user,
	}
	token, err := jwtToken.SignedString([]byte(os.Getenv("APP_KEY")))
	return value, token, err
}

func (uts eventHubUserTokenService) GenerateToken(user *models.EventHubUser) (uint64, string, error) {
	return uts.generateJwtToken(user)
}

func (uts eventHubUserTokenService) UpdateUserTokenInDB(userID uint64, time time.Time, token string) error {
	userToken := repositories.EventHubUserTokenRepository.GetUserTokenByUserId(userID)
	if userToken == nil {
		return errors.New("No user does not have user_token")
	}
	userToken.Token = token
	userToken.ExpiredAt = time

	return repositories.EventHubUserTokenRepository.Update(userToken)
}

func (uts eventHubUserTokenService) CreateUserTokenInDB(user *models.EventHubUser) error {
	_, token, _ := uts.generateJwtToken(user)
	return repositories.EventHubUserTokenRepository.Create(&models.EventHubUserToken{
		UserID: user.Id,
		Token:  token,
	})
}

func (uts eventHubUserTokenService) CreateOrUpdateUserTokenInDB(user *models.EventHubUser) error {
	_, token, _ := uts.generateJwtToken(user)

	dbToken, errs := repositories.EventHubUserTokenRepository.GetUserTokenByUserIdOnCreate(user.Id)
	if errs != nil {
		return repositories.EventHubUserTokenRepository.Create(&models.EventHubUserToken{
			UserID: user.Id,
			Token:  token,
		})
	}
	dbToken.Token = token
	return repositories.EventHubUserTokenRepository.Update(dbToken)

}

func (uts eventHubUserTokenService) GetUserClaimFromLocal(ctx *fiber.Ctx) *models.EventHubUserClaims {
	return ctx.Locals("user").(*models.EventHubUserClaims)
}

func (uts eventHubUserTokenService) GetUserFromLocal(ctx *fiber.Ctx) *models.EventHubUser {
	return uts.GetUserClaimFromLocal(ctx).EventHubUser
}
