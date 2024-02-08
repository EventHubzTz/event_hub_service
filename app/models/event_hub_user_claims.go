package models

import "github.com/golang-jwt/jwt"

type EventHubUserClaims struct {
	*jwt.StandardClaims
	TokenType string
	*EventHubUser
}
