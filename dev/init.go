package dev

import (
	"log"

	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/structs"
)

func Init() {
	log.Println("Starting with development mode.")

	var profile structs.Profile
	res := db.Core.First(&profile)
	if res.Error != nil {
		CreateTestEntry()
	}

	log.Println("Discord credentials: ", config.Conf.Credentials)
	log.Println("Callback URL: ", config.Conf.General.CallbackURL)
}
