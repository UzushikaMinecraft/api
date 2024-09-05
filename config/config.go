package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/uzushikaminecraft/api/structs"
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
