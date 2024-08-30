package main

import (
	"log"

	"github.com/2mugi/uzsk-api/api"
	"github.com/2mugi/uzsk-api/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init config
	Conf := config.Init()

	// Init web server
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("200 OK")
	})

	// API routes
	// - /api/server
	app.Get("/api/servers/:name", func(c *fiber.Ctx) error {
		return c.JSON(api.GetServer(Conf, c.Params("name")))
	})

	log.Fatal(app.Listen(":3000"))
}
