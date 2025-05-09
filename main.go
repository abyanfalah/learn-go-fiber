package main

import (
	"learn-fiber/core/authutil"
	"learn-fiber/core/config/database"
	"learn-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "asdf",
		AppName:       "asdf",
		// ErrorHandler:  core.ErrorHandler,
	})

	// gormDB, sqlDB := database.Connect()
	database.Connect()
	database.InitMigration(database.DB)

	router.SetupRouters(app)

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(authutil.JwtConfig())

	// app.Use(logger.New(logger.Config{
	// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// }))

	app.Use(logger.New())

	log.Fatal(app.Listen(":3000"))
}
