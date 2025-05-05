package auth

import "github.com/gofiber/fiber/v2"

func router(app *fiber.App) {
	api := app.Group("/auth")

	api.Get("/login", login)
}
