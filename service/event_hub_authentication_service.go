package service

import (
	"golang.org/x/crypto/bcrypt"
)

var EventHubAuthenticationService = newEventHubAuthenticationService()

type eventHubAuthenticationService struct {
}

func newEventHubAuthenticationService() eventHubAuthenticationService {
	return eventHubAuthenticationService{}
}

func (_ eventHubAuthenticationService) HashPassword(password string) string {
	byt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(byt)
}

func (_ eventHubAuthenticationService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
