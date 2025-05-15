package helper

import (
	"fmt"
	"learn-fiber/core/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseAndValidate[T any](c *fiber.Ctx) (*T, *exception.ErrValidation) {
	req := new(T)

	if err := c.BodyParser(&req); err != nil {
		return nil, &exception.ErrValidation{
			Message: fmt.Sprintf("Invalid payload: %s", err.Error()),
		}
	}

	// Validate the struct
	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string]any)

		for _, err := range err.(validator.ValidationErrors) {
			validationMsg := strings.TrimSpace(
				fmt.Sprintf("%s %s",
					err.ActualTag(),
					err.Param()))
			validationErrors[err.Field()] = validationMsg
		}

		return nil, &exception.ErrValidation{
			Message:     "Struct validation failed",
			Validations: validationErrors,
		}
	}

	return req, nil
}
