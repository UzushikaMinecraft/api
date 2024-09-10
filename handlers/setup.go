package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/uzushikaminecraft/api/config"
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
func Setup() {
	app := fiber.New()

	// API routes
	api := app.Group("/api")

	// /api/servers
	servers := api.Group("/servers")
	servers.Get("/:name", HandleServer)
	servers.Get("/", HandleServers)

	// /api/profiles
	profiles := api.Group("/profiles")
	profiles.Get("/:uuid", HandleProfile)
	profiles.Get("/", HandleProfiles)

	// /api/auth
	auth := api.Group("/auth")
	auth.Get("/", HandleAuth)
	auth.Get("/callback", HandleAuthCallback)

	// /api/me
	me := api.Group("/me")
	me.Get("/", HandleMe)
	me.Post("/biography", HandleMeBiography)

	// /api/discord
	discord := api.Group("/discord")
	discord.Get("/:uuid", HandleDiscordUuid)

	// /api/avatar
	avatar := api.Group("/avatar")
	avatar.Get("/:part/bedrock/:xuid", HandleRenderBedrockSkin)

	// Swagger
	// /api/swagger
	api.Get("/swagger/*", swagger.HandlerDefault)

	// Logger
	app.Use(logger.New())

	// Encrypt cookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.Conf.Credentials.JWTSecret,
	}))

	// Run the web server
	log.Fatal(app.Listen(":3200"))
}
