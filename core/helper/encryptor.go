package helper

import (
	"fmt"
	"learn-fiber/core/exception"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var encryptionCost = 10

func GetEncryptedPassword(password string) (string, error) {
	password = strings.TrimSpace(password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), encryptionCost)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "error on password encrypting: "+err.Error())
	}

	return string(hashedPassword), nil
}

func IsCorrectPassword(hashedPassword string, rawPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		fmt.Println("error comparing password: ", err)
		return false, exception.ErrBadCredentials
	}

	return true, nil
}
