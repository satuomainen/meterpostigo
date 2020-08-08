package db

import (
	"com.github/satuomainen/meterpostigo/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", // user:pass@(host:port)/dbname
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)

	var err error
	DB, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       dataSourceName,
			SkipInitializeWithVersion: false,
		}),
		&gorm.Config{})


	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database, exiting - %s", err))
	}
}
