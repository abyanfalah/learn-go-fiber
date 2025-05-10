package mdw

import (
	"errors"
	"fmt"
	"learn-fiber/core/config"
	"learn-fiber/core/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

var isWhiteListedPath = map[string]bool{
	"/":                  true,
	"/api":               true,
	"/api/auth/login":    true,
	"/api/auth/register": true,
}

func AuthCookieMiddleware() fiber.Handler {
	// hehe
	return logic
}

func logic(c *fiber.Ctx) error {
	if isWhiteListedPath[c.Path()] {
		return c.Next()
	}

	token, err := parseJWTFromCookie(c, config.GetEnv("JWT_SECRET"))
	if err != nil {
		log.Error(err.Error())
		return exception.ErrUnauthorized
	}

	claims, err := validateToken(token)
	if err != nil {
		log.Error(err.Error())
		return exception.ErrUnauthorized
	}

	// dont know this yet
	c.Locals("user", claims) // or token if you want raw access
	return c.Next()
}

func parseJWTFromCookie(c *fiber.Ctx, secret string) (*jwt.Token, error) {
	tokenString := c.Cookies(config.GetEnv("COOKIE_NAME"))
	fmt.Println("token -> " + tokenString)
	if tokenString == "" {
		return nil, errors.New("missing token")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
}

func validateToken(token *jwt.Token) (jwt.MapClaims, error) {
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	fmt.Println("token is valid")
	fmt.Println("claims: ")
	fmt.Println(claims)

	return claims, nil
}
