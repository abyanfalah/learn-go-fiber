package exception

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	ErrBadCredentials = fiber.NewError(fiber.StatusBadRequest, "bad credentials")
)

// Handle i wish i can catch an exception from the sky
func Handle(e error) *fiber.Error {
	if e == nil {
		return nil
	}

	switch {
	case errors.Is(e, gorm.ErrRecordNotFound):
		return fiber.NewError(fiber.StatusNotFound, "resource not found")

	default:
		return fiber.NewError(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR: "+e.Error())
	}

}

func BadCredentials(e error) *fiber.Error {
	return fiber.NewError(fiber.StatusBadRequest, "bad credentials")
}
