package response

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
)

type ErrorResponseStruct struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

type ErrorResponseStructString struct {
	Error   bool     `json:"error"`
	Message string `json:"message"`
}

type successResponseStruct struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type externalServiceDataListResponseStruct struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

type externalServiceDataListSuccessResponseStruct struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type successResponseWithUserTokenStruct struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    uint64 `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ErrorResponse(message string, code int, ctx *fiber.Ctx) error {
	var messages []string
	return ctx.Status(code).JSON(ErrorResponseStruct{
		Message: append(messages, message),
		Error:   true,
	})
}
func ErrorResponseStr(message string, code int, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(ErrorResponseStructString{
		Message: message,
		Error:   true,
	})
}

func SuccessResponse(message string, code int, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(successResponseStruct{
		Message: message,
		Error:   false,
	})
}

func DataListResponse(response interface{}, code int, ctx *fiber.Ctx) error {
	modelResponse := reflect.ValueOf(response)
	return ctx.Status(code).JSON(externalServiceDataListResponseStruct{
		Data:  response,
		Count: modelResponse.Len(),
	})
}

func DataListSuccessResponse(response interface{}, code int, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(externalServiceDataListSuccessResponseStruct{
		Data:  response,
		Error: false,
	})
}

func MapDataResponse(response interface{}, code int, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(response)
}

func SuccessResponseWithUserToken(token string, code int, expiresIn uint64, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(successResponseWithUserTokenStruct{
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
		AccessToken:  token,
		RefreshToken: "",
	})
}

func InternalServiceDataResponse(response interface{}, code int, ctx *fiber.Ctx) error {
	return ctx.Status(code).JSON(response)
}