package user

import (
	"learn-fiber/core/config/database"
	"learn-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func isUsedEmail(email string) error {
	var isExists bool
	err := database.DB.Model(&model.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&isExists).Error

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if isExists {
		return fiber.NewError(fiber.StatusBadRequest, "email is used")
	}

	return nil
}
