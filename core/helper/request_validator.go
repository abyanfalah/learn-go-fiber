package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate[T any](c *fiber.Ctx) (*T, error) {
	req := new(T)

	if err := c.BodyParser(&req); err != nil {
		fmt.Println("error parsing")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}

	if err := validate.Struct(req); err != nil {
		fmt.Println("error struct validation")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Struct validation failed")
	}

	return req, nil
}
