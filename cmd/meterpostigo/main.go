package main

import (
	"com.github/satuomainen/meterpostigo/internal/config"
	"fmt"
)

func main() {
	fmt.Println("Database host:", config.Config.Database.Host)
	fmt.Println("Database password:", config.Config.Database.Password)
}
