package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/dev"
	_ "github.com/uzushikaminecraft/api/docs"
	"github.com/uzushikaminecraft/api/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	Conf := config.Init(confPath)

	// Init db
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		Conf.MySQL.User,
		Conf.MySQL.Password,
		Conf.MySQL.Host,
		Conf.MySQL.Port,
		Conf.MySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	db.AutoMigrate(&structs.Profile{})

	// For development
	if strings.Contains(os.Args[0], "go-build") {
		dev.CreateTestEntry(db)
	}

	// Init web server
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("200 OK")
	})

	// - /api/servers
	app.Get("/api/servers", func(c *fiber.Ctx) error {
		return c.JSON(services.GetServers(Conf))
	})

	// API routes
	// - /api/server
	app.Get("/api/servers/:name", func(c *fiber.Ctx) error {
		return c.JSON(services.GetServer(Conf, c.Params("name")))
	})

	// Player
	// - /api/profiles
	app.Get("/api/profiles", func(c *fiber.Ctx) error {
		m := c.Queries()
		return c.JSON(services.GetProfiles(db, m))
	})

	// - /api/profiles/:uuid
	app.Get("/api/profiles/:uuid", func(c *fiber.Ctx) error {
		return c.JSON(services.GetProfile(db, c.Params("uuid")))
	})

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	log.Fatal(app.Listen(":3000"))
}
