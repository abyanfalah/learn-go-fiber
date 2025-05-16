package general

import (
	"learn-fiber/core/mdw"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r fiber.Router) {
	r.Get("", test)
	r.Get("error", testError)
	r.Get("protected", mdw.AuthCookieMiddleware(), testProtected)
	r.Post("", testWithPayload)
	r.Get("user-dummy", createUser)
}
