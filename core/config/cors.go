package config

import (
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var originsToAllow = []string{
	"*",
}

func CorsConfig() cors.Config {
	allowedOrigins := strings.Join(originsToAllow, ", ")

	return cors.Config{
		AllowOrigins: allowedOrigins,
		// AllowOrigins: "http://www.google.com",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
}
