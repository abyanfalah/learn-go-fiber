package main

import (
	"learn-fiber/core/config"
	"learn-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitEnv()
	app := fiber.New()

	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
	// fmt.Print(config.GetEnv("APP_ENV"))

}
