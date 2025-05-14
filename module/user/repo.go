package user

import (
	"learn-fiber/core/config/db"
	"learn-fiber/model"
	"learn-fiber/module/auth"

	"github.com/gofiber/fiber/v2"
)

func IsUsedEmail(email string) error {
	var isExists bool
	err := db.Use().Model(&model.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&isExists).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if isExists {
		return fiber.NewError(fiber.StatusBadRequest, "email is used")
	}

	auth.Hehe()

	return nil
}
