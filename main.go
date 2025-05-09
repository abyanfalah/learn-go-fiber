package main

import (
	"learn-fiber/core/authutil"
	"learn-fiber/core/config"
	"learn-fiber/core/config/database"
	"learn-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// defer database.DB.clo

	// app := fiber.New(fiber.Config{
	// 	CaseSensitive: true,
	// 	StrictRouting: true,
	// 	ServerHeader:  "asdf",
	// 	AppName:       "asdf",
	// 	// ErrorHandler:  core.ErrorHandler,
	// })

	app := fiber.New()
	app.Use(logger.New())

	database.Connect()
	database.InitMigration(database.DB)

	router.SetupRouters(app)

	app.Use(cors.New(config.CorsConfig()))
	app.Use(authutil.JwtConfig())

	// app.Use(recover.New())

	log.Fatal(app.Listen(":3000"))
}
