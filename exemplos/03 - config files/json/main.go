package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	LoadConfigFromJson()
}

func LoadConfigFromJson() {

	fmt.Println("Reading config from JSON file...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	port := viper.Get("database.port")
	host := viper.Get("database.host")
	name := viper.Get("database.name")

	fmt.Println(host, port, name)
}
