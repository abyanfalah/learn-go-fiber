package auth

import "github.com/gofiber/fiber/v2"

func SetupRoutes(r fiber.Router) {
	api := r.Group("auth")

	api.Post("login", login)
	api.Post("register", register)
	api.Post("logout", logout)
}
