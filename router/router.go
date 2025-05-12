package router

import (
	"learn-fiber/core/mdw"
	"learn-fiber/module/auth"
	"learn-fiber/module/general"
	"learn-fiber/module/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	api := app.Group("api", mdw.AuthCookieMiddleware())

	general.SetupRoutes(api)
	auth.SetupRoutes(api)
	user.SetupRoutes(api)
}
