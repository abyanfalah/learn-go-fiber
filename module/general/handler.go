package general

import (
	"learn-fiber/core/http/response"

	"github.com/gofiber/fiber/v2"
)

type orang struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test(c *fiber.Ctx) error {
	return response.Success(c)
}

func testWithPayload(c *fiber.Ctx) error {
	o := new(orang)
	o.Age = 10
	o.Name = "ohang"

	return response.Body(c, o)
}
