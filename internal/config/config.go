package config

import "fmt"

func GetStringValue(name string) (value string, err error) {
	value = fmt.Sprintf("Value of %s", name)
	return value, nil
}
