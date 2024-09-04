package config

import (
	"log"

	"github.com/uzushikaminecraft/uzsk-api/structs"
	"github.com/BurntSushi/toml"
)

func Init(path string) structs.Config {
	// Init configuration
	var Conf structs.Config
	_, err := toml.DecodeFile(path, &Conf)

	if err != nil {
		log.Fatalln(err)
	}

	return Conf
}
