package auth

import "github.com/gofiber/fiber/v2"

func login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "login",
	})
}
