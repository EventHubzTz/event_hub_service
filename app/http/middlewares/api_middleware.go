package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var HomeToken string

func Api(ctx *fiber.Ctx) error {
	if !strings.Contains(string(ctx.Request().Header.ContentType()), "application/json") && !strings.Contains(string(ctx.Request().Header.ContentType()), "multipart/form-data") {
		return ctx.SendString("Bad request")
	}

	HomeToken = ctx.Get("event_hub_auth")

	return ctx.Next()
}
