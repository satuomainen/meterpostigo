package main

import (
	"com.github/satuomainen/meterpostigo/internal/config"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	configValue, err := config.GetStringValue("Any")
	if err != nil {
		panic("Configuration error, cannot operate.")
	}

	fmt.Println("Configuration today is:", configValue)
}
