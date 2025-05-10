package main

import (
	"learn-fiber/core/authutil"
	"learn-fiber/core/config"
	"learn-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))

	router.SetupRouters(app)
	app.Use(authutil.JwtConfig())

	log.Fatal(app.Listen(":3000"))
}
