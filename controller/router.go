package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("200 OK")
	})

	// API routes
	// Retrieve all servers
	app.Get("/api/servers", func(c *fiber.Ctx) error {
		return c.JSON(services.GetServers())
	})

	// Retrieve specified server's information
	app.Get("/api/servers/:name", func(c *fiber.Ctx) error {
		res, _ := services.GetServer(c.Params("name"))
		return c.JSON(res)
	})

	// Retrieve all players
	app.Get("/api/profiles", func(c *fiber.Ctx) error {
		m := c.Queries()

		res, err := services.GetProfiles(m)
		if err != nil {
			return c.Status(500).JSON(
				structs.Error{
					Error: err.Error(),
				},
			)
		}

		return c.JSON(res)
	})

	// Retrieve players' profile from UUID
	app.Get("/api/profiles/:uuid", func(c *fiber.Ctx) error {
		res, _ := services.GetProfile(c.Params("uuid"))
		return c.JSON(res)
	})

	// Retrieve Bedrock players' skin from provided XUID
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

	// OAuth callback endpoint
	app.Get("/api/login", auth.Login)
	app.Get("/api/login/callback", auth.Callback)

	// Personal information
	app.Get("/api/me", func(c *fiber.Ctx) error {
		res, err := services.GetMe(c.Get("X-Auth-Token"))

		if err == nil {
			return c.Status(200).JSON(
				res,
			)
		}

		if err.Error() == "token is not provided" || err.Error() == "invalid token" {
			return c.Status(400).JSON(
				structs.Error{
					Error: err.Error(),
				},
			)
		}

		return c.Status(500).JSON(
			structs.Error{
				Error: err.Error(),
			},
		)
	})

	// Retrieve UUID from provided Discord ID
	app.Get("/api/discord/:uuid", services.GetUUIDByDiscord)

	// Swagger
	app.Get("/api/swagger/*", swagger.HandlerDefault)
}
