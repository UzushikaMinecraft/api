package config

import (
	"log"

	"github.com/2mugi/uzsk-api/structs"
	"github.com/BurntSushi/toml"
)

func Init() structs.Config {
	// Init configuration
	var Conf structs.Config
	_, err := toml.DecodeFile("./config.toml", &Conf)

	if err != nil {
		log.Fatalln(err)
	}

	return Conf
}
