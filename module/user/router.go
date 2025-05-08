package user

import "github.com/gofiber/fiber/v2"

func SetupRoutes(r fiber.Router) {
	api := r.Group("user")

	api.Get("", getAll)
	api.Post("", create)
	api.Put("/:id", update)
	api.Delete("/:id", delete)
}
