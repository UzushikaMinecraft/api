package dev

import (
	"log"

	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/structs"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	log.Println("Starting with development mode.")

	var profile structs.Profile
	res := db.First(&profile)
	if res.Error != nil {
		CreateTestEntry(db)
	}

	log.Println("Discord credentials: ", config.Conf.Credentials)
	log.Println("Callback URL: ", config.Conf.General.CallbackURL)
}
