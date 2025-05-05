package response

import "github.com/gofiber/fiber/v2"

// SuccessWithMessage self explanatory
func SuccessWithMessage(c *fiber.Ctx, msg string) error {
	return c.JSON(fiber.Map{
		"message": msg,
	})
}

// Success response
func Success(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}

// Body sends json
func Body(c *fiber.Ctx, data any) error {
	return c.JSON(data)
}

// Body sends json
func BodyWithMessage(c *fiber.Ctx, data any) error {
	return c.Send(c.Body())
}
