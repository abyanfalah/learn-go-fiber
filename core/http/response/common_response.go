package response

import (
	"learn-fiber/core/exception"

	"github.com/gofiber/fiber/v2"
)

func SuccessWithMessage(c *fiber.Ctx, msg string) error {
	return c.JSON(fiber.Map{
		"message": msg,
	})
}

func Success(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Body(c *fiber.Ctx, data any) error {
	return c.JSON(data)
}

func InvalidPayload(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid payload",
	})
}

func InvalidCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Invalid credentials",
	})
}

func ErrorValidation(c *fiber.Ctx, ev *exception.ErrValidation) error {
	return c.Status(fiber.StatusBadRequest).JSON(ev)
}
