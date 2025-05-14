package auth

import (
	"learn-fiber/core/authutil"
	"learn-fiber/core/exception"
	"learn-fiber/core/helper"
	"learn-fiber/core/helper/generator"
	"learn-fiber/core/http/response"
	baseresponse "learn-fiber/http/base_response"
	"learn-fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	req, _ := helper.ParseAndValidate[loginRequest](c)

	user, err := findUserByEmailAndPassword(req)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	token, err := authutil.SetCookie(c, user)
	if err != nil {
		return exception.Handle(err)
	}

	r := baseresponse.AuthResponse{
		AccessToken: token,
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
	}

	return response.Body(c, r)
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

	// err = isUsedEmail(u.Email)

	return response.Success(c)
}

func logout(c *fiber.Ctx) error {
	err := authutil.BlackListToken(authutil.GetJwt(c), time.Hour*12)
	if err != nil {
		return exception.Handle(err)
	}

	authutil.ClearCookie(c)
	return response.Success(c)
}
