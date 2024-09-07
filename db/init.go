package db

import (
	"fmt"

	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.User,
		config.Conf.MySQL.Password,
		config.Conf.MySQL.Host,
		config.Conf.MySQL.Port,
		config.Conf.MySQL.Database,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&structs.Profile{})

	return nil
}
