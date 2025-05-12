package general

import (
	"errors"
	"learn-fiber/core/config/database"
	"learn-fiber/core/helper/generator"
	"learn-fiber/core/http/response"
	"learn-fiber/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orang struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test(c *fiber.Ctx) error {
	err := errors.New("asdf")
	return err
}

func testProtected(c *fiber.Ctx) error {
	return c.JSON(c.Locals("user"))
}

func testWithPayload(c *fiber.Ctx) error {
	o := new(orang)
	o.Age = 10
	o.Name = "ohang"

	return response.Body(c, o)
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

	database.DB.Save(&u)

	return response.Success(c)
}
