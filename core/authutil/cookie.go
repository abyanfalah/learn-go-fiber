package authutil

import (
	"learn-fiber/core/config"
	"learn-fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

var isCookieSecure = false

func SetCookie(c *fiber.Ctx, user *model.User) (string, error) {
	token, err := createJwt(user)
	if err != nil {
		return "", err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = config.GetEnv("COOKIE_NAME")
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Secure = isCookieSecure
	cookie.HTTPOnly = true

	c.Cookie(cookie)

	return token, nil
}

// c.clearCookie() is not working. might figure out later
func ClearCookie(c *fiber.Ctx) {
	cookie := new(fiber.Cookie)
	cookie.Name = config.GetEnv("COOKIE_NAME")
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Second)
	cookie.Secure = isCookieSecure
	cookie.HTTPOnly = true

	c.Cookie(cookie)
}
