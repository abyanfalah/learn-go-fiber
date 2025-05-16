package mdw

import (
	"errors"
	"learn-fiber/core/authutil"
	"learn-fiber/core/config"
	"learn-fiber/core/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

var isWhiteListedPath = map[string]bool{
	"/":                  true,
	"/api":               true,
	"/api/error":         true,
	"/api/auth/login":    true,
	"/api/auth/register": true,
}

func AuthCookieMiddleware() fiber.Handler {
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

	if token == nil {
		log.Error("missing token")
		return exception.ErrInvalidToken
	}

	isBlacklisted, err := authutil.IsBlacklistedToken(token.Raw)
	if err != nil {
		log.Error(err.Error())
		return exception.Handle(err)
	}

	if isBlacklisted {
		log.Error("token is blacklisted")
		return exception.Unauthorized("token is blacklisted")
	}

	claims, err := validateToken(token)
	if err != nil {
		log.Error(err.Error())
		return exception.ErrUnauthorized
	}

	c.Locals("user", claims)
	return c.Next()
}

func parseJWTFromCookie(c *fiber.Ctx, secret string) (*jwt.Token, error) {
	tokenString := c.Cookies(config.GetEnv("COOKIE_NAME"))
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
	return claims, nil
}
