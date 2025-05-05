package general

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r fiber.Router) {
	r.Get("/", test)
	r.Post("/", testWithPayload)
}
