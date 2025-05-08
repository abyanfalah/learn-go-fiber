package mdw

import (
	"github.com/gofiber/fiber/v2"
)

var isWhiteListedPath = map[string]bool{
	"/api/auth/login":    true,
	"/api/auth/register": true,
}

func JwtAuth(c *fiber.Ctx) error {
	if isWhiteListedPath[c.Path()] {
		return c.Next()
	}

	// fmt.Println(c.coo

	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)

	// fmt.Println("you are: " + name)

	return c.Next()
}
