package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/dev"
	_ "github.com/uzushikaminecraft/api/docs"
	"github.com/uzushikaminecraft/api/login"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @title uzsk-api
// @version 1.0
// @description Public Web API for uzsk.iamtakagi.net
// @termsOfService https://uzsk.iamtakagi.net
// @contact.name yude
// @contact.email i@yude.jp
// @license.name MIT
// @license.url https://opensource.org/license/mit
// @host uzsk-api.iamtakagi.net
// @BasePath /api
func main() {
	// Flag definitions
	confPath := ""
	flag.StringVar(&confPath, "config", "./config.toml", "path to config.toml")
	flag.Parse()

	// Init config
	err := config.Init(confPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Init db
	err = db.Init()
	if err != nil {
		log.Fatalln(err)
	}

	// Init Discord OAuth
	login.Init()

	// For development
	if strings.Contains(os.Args[0], "go-build") {
		dev.Init(db.DB)
	}

	// Init web server
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("200 OK")
	})

	// API routes
	// - /api/servers
	app.Get("/api/servers", func(c *fiber.Ctx) error {
		return c.JSON(services.GetServers())
	})

	// - /api/server
	app.Get("/api/servers/:name", func(c *fiber.Ctx) error {
		res, _ := services.GetServer(c.Params("name"))
		return c.JSON(res)
	})

	// Player
	// - /api/profiles
	app.Get("/api/profiles", func(c *fiber.Ctx) error {
		m := c.Queries()

		res, err := services.GetProfiles(db.DB, m)
		if err != nil {
			return c.Status(500).JSON(
				structs.Error{
					Error: err.Error(),
				},
			)
		}

		return c.JSON(res)
	})

	// - /api/profiles/:uuid
	app.Get("/api/profiles/:uuid", func(c *fiber.Ctx) error {
		res, _ := services.GetProfile(db.DB, c.Params("uuid"))
		return c.JSON(res)
	})

	// - /api/avatar/:part/bedrock/:xuid
	app.Get("/api/avatar/:part/bedrock/:xuid", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "image/png")

		img, err := services.RenderBedrockSkin(c.Params("xuid"), c.Params("part"))
		if err != nil {
			return c.Status(500).JSON(structs.Error{
				Error: "Error occured while rendering image",
			})
		}

		c.Response().Header.Set("Cache-Control", "public, max-age=86400, immutable")
		return c.SendStream(img)
	})

	// Swagger
	app.Get("/api/swagger/*", swagger.HandlerDefault)

	// CORS settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// OAuth callback endpoint
	app.Get("/api/login", login.Login)
	app.Get("/api/login/callback", login.Callback)

	// Run the web server
	log.Fatal(app.Listen(":3000"))
}
