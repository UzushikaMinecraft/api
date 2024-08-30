package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/2mugi/uzsk-api/api"
	"github.com/2mugi/uzsk-api/config"
	"github.com/2mugi/uzsk-api/dev"
	"github.com/2mugi/uzsk-api/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Init config
	Conf := config.Init()

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
	db.AutoMigrate(&structs.Player{})

	// For development
	if strings.Contains(os.Args[0], "go-build") {
		dev.CreateTestEntry(db)
	}

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

	// Player
	// - /api/players
	app.Get("/api/players", func(c *fiber.Ctx) error {
		m := c.Queries()
		return c.JSON(api.GetPlayers(db, m))
	})

	log.Fatal(app.Listen(":3000"))
}
