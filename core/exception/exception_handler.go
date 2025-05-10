package exception

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func Handler(c *fiber.Ctx, err error) error {
// 	return Handle(err)
// }

func Handle(e error) *fiber.Error {
	if e == nil {
		return nil
	}

	logError(e)

	switch {
	case errors.Is(e, gorm.ErrRecordNotFound):
		return fiber.NewError(fiber.StatusNotFound, "resource not found")

	default:
		return fiber.NewError(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR: "+e.Error())
	}

}
