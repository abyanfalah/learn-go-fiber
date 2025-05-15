package user

import (
	"learn-fiber/core/config/db"
	"learn-fiber/core/helper"
	"learn-fiber/core/http/response"
	"learn-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func getAll(c *fiber.Ctx) error {
	var users []model.User
	db.Use().Find(&users)

	return response.Body(c, users)
}

func create(c *fiber.Ctx) error {
	req, ve := helper.ParseAndValidate[createRequest](c)
	if ve != nil {
		return ve
	}

	err := IsUsedEmail(req.Email)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	hashedPassword, err := helper.GetEncryptedPassword(req.Password)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	u := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	db.Use().Save(&u)

	return response.Success(c)
}

func update(c *fiber.Ctx) error {
	req, err := helper.ParseAndValidate[updateRequest](c)
	if err != nil {
		return err
	}

	var id int = helper.ToInt(c.Params("id"))
	var user model.User

	result := db.Use().First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	if user.Email != req.Email {
		err := IsUsedEmail(req.Email)
		if err != nil {
			return fiber.DefaultErrorHandler(c, err)
		}
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	db.Use().Save(&user)
	return response.Body(c, user)
}

func delete(c *fiber.Ctx) error {
	var id int = helper.ToInt(c.Params("id"))
	var user model.User

	result := db.Use().First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	db.Use().Delete(&user)
	return response.Success(c)
}
