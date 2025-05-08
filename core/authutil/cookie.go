package authutil

import (
	"learn-fiber/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var isCookieSecure = false

func SetCookie(c *fiber.Ctx, user *model.User) error {
	token, err := createJwt(user)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	// cookie.Expires = time.Now().Add(5 * time.Second)
	// cookie.Expires = time.Now().Add(1 * time.Minute)
	cookie.Secure = isCookieSecure
	cookie.HTTPOnly = true

	c.Cookie(cookie)

	return nil
}

// c.clearCookie() is not working. might figure out later
func ClearCookie(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Second)
	cookie.Secure = isCookieSecure
	cookie.HTTPOnly = true

	c.Cookie(cookie)
	return nil
}

func createJwt(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("kunyuk444"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
