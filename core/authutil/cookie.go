package authutil

import (
	"learn-fiber/core/config"
	"learn-fiber/model"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
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
	cookie.Name = config.GetEnv("COOKIE_NAME")
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Secure = isCookieSecure
	cookie.HTTPOnly = true

	c.Cookie(cookie)

	return nil
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

func createJwt(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JwtConfig() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET"))},
	})
}
