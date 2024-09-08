package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/controller"
	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/dev"
	_ "github.com/uzushikaminecraft/api/docs"
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
	auth.Init()

	// For development
	if strings.Contains(os.Args[0], "go-build") {
		dev.Init()
	}

	// Init web server
	app := fiber.New()

	// Setup routes
	controller.SetupRoutes(app)

	// CORS settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Run the web server
	log.Fatal(app.Listen(":3000"))
}
