package config

import (
	"github.com/BurntSushi/toml"
	"github.com/uzushikaminecraft/api/structs"
)

var Conf structs.Config

func Init(path string) error {
	// Init configuration
	_, err := toml.DecodeFile(path, &Conf)

	if err != nil {
		return err
	}

	return nil
}
