package auth

import (
	"learn-fiber/core/authutil"
	"learn-fiber/core/config/database"
	"learn-fiber/core/exception"
	"learn-fiber/core/helper"
	"learn-fiber/core/helper/generator"
	"learn-fiber/core/http/response"
	"learn-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	req, _ := helper.ParseAndValidate[loginRequest](c)

	user, err := findUserByEmailAndPassword(req)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	err = authutil.SetCookie(c, user)
	if err != nil {
		return exception.Handle(err)
	}

	return response.Body(c, user)
}

func register(c *fiber.Ctx) error {
	req, err := helper.ParseAndValidate[registerRequest](c)
	if err != nil {
		return err
	}

	u := new(model.User)
	u.ID = generator.GenerateId()
	u.Name = req.Name
	u.Email = req.Email
	u.Password = req.Password

	database.DB.Save(u)
	return response.Success(c)
}

func logout(c *fiber.Ctx) error {
	err := authutil.ClearCookie(c)
	if err != nil {
		return exception.Handle(err)
	}
	return response.Success(c)
}
