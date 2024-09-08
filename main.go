package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/dev"
	_ "github.com/uzushikaminecraft/api/docs"
	"github.com/uzushikaminecraft/api/handlers"
)

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

	// Init / run the web server
	handlers.Setup()
}
