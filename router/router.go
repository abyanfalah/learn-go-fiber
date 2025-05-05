package router

import (
	"learn-fiber/module/general"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")

	general.SetupRoutes(api)

}
