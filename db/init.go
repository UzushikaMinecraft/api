package db

import (
	"fmt"

	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Core *gorm.DB
var DiscordSRV *gorm.DB

func Init() error {
	var err error

	dsnCore := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.Core.User,
		config.Conf.MySQL.Core.Password,
		config.Conf.MySQL.Core.Host,
		config.Conf.MySQL.Core.Port,
		config.Conf.MySQL.Core.Database,
	)

	dsnDiscordSRV := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.DiscordSRV.User,
		config.Conf.MySQL.DiscordSRV.Password,
		config.Conf.MySQL.DiscordSRV.Host,
		config.Conf.MySQL.DiscordSRV.Port,
		config.Conf.MySQL.DiscordSRV.Database,
	)

	Core, err = gorm.Open(mysql.Open(dsnCore), &gorm.Config{})
	if err != nil {
		return err
	}

	DiscordSRV, err = gorm.Open(mysql.Open(dsnDiscordSRV), &gorm.Config{})
	if err != nil {
		return err
	}

	Core.AutoMigrate(&structs.Profile{})
	DiscordSRV.AutoMigrate(&structs.DiscordSrvAccounts{})

	return nil
}
