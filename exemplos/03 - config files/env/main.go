package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func main() {
	LoadConfigFromEnv()
	LoadEnvironmentVariables()

}

// reference: https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d
func LoadConfigFromEnv() {
	var config Config

	fmt.Println("Reading config from config.env...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config.DBDriver)
	fmt.Println(config.DBSource)
	fmt.Println(config.ServerAddress)

}

/*
CMD (Windows)

  - Abra o terminal (cmd)

  - crie as variaveis de ambiente

    set DB_USERNAME=admin
    set DB_PASSWORD=123456

PowerShell (Windows)

	$env:DB_USERNAME="admin"
	$env:DB_PASSWORD

Linux

	export DB_USERNAME=admin
	export DB_PASSWORD=123456
*/
func LoadEnvironmentVariables() {

	fmt.Println("\nReading environment variables...")

	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")

	databaseUsername := viper.GetString("database.username")
	databasePassword := viper.GetString("database.password")

	fmt.Println("Database Username from Environment Variable:", databaseUsername)
	fmt.Println("Database Password from Environment Variable:", databasePassword)

}
