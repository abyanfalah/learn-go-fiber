package exception

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrBadCredentials = fiber.NewError(fiber.StatusBadRequest, "bad credentials")
	ErrUnauthorized   = fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	ErrInvalidToken   = fiber.NewError(fiber.StatusUnauthorized, "invalid token")
)

func BadCredentials(e error) *fiber.Error {
	fmt.Println("bad credentials: " + e.Error())
	return ErrBadCredentials
}

func Unauthorized(msg string) *fiber.Error {
	return fiber.NewError(fiber.StatusUnauthorized, msg)
}
