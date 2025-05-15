package general

import (
	"learn-fiber/core/config/db"
	"learn-fiber/core/helper"
	"learn-fiber/core/helper/generator"
	"learn-fiber/core/http/response"
	"learn-fiber/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// type orang struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

func test(c *fiber.Ctx) error {
	return response.Success(c)
}

func testProtected(c *fiber.Ctx) error {
	return c.JSON(c.Locals("user"))
}

func testWithPayload(c *fiber.Ctx) error {
	req, ev := helper.ParseAndValidate[testRequest](c)
	if ev != nil {
		return response.ErrorValidation(c, ev)
	}

	return response.Body(c, req)
}

func createUser(c *fiber.Ctx) error {

	id := generator.GenerateId()
	name := strconv.Itoa(int(id))
	email := name + "@gmail.com"

	u := model.User{}
	u.ID = id
	u.Name = name
	u.Email = email
	u.Password = name

	db.Use().Save(&u)

	return response.Success(c)
}
