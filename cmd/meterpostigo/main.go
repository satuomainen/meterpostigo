package main

import (
	"com.github/satuomainen/meterpostigo/internal/config"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/model"
	"fmt"
)

func main() {
	fmt.Println("Database host    :", config.Config.Database.Host)
	fmt.Println("Database password:", config.Config.Database.Password)

	var dataseries = new(model.DataSeries)
	db.DB.First(&dataseries)

	fmt.Printf("Sign of life from the database: %+v\n", dataseries)
}
