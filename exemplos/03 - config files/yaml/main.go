package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type NestedURL struct {
	URL    string `mapstructure:"url"`
	Nested int    `mapstructure:"nested"`
}

type ZipCode struct {
	Zipcode string `mapstructure:"zipcode"`
}

type Config struct {
	Element []map[string]NestedURL `mapstructure:"element"`
	Weather []map[string]ZipCode   `mapstructure:"weather"`
}

func main() {
	LoadConfigFromYaml()

	fmt.Println("\n--------------------------------------------------")

	LoadConfigFromYaml2()

	WriteValuesToConfigFile()
}

// reference:  https://stackoverflow.com/questions/52585309/how-do-i-use-viper-to-get-a-value-from-a-nested-yaml-structure
func LoadConfigFromYaml() {

	fmt.Println("Reading config from config.yaml...")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)

	for _, element := range config.Element {
		for key, value := range element {
			fmt.Println(key)
			fmt.Println(value)
		}

	}

	for _, weather := range config.Weather {
		for key, value := range weather {
			fmt.Println(key)
			fmt.Println(value)
		}
	}

}

// reference: https://reintech.io/blog/managing-configuration-go-application-viper-tutorial
func LoadConfigFromYaml2() {

	viper.AddConfigPath(".")
	viper.SetConfigName("config2")
	viper.SetConfigType("yaml")

	fmt.Println("Reading config from config.yaml...")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	serverHost := viper.GetString("server.host")
	serverPort := viper.GetInt("server.port")
	databaseUsername := viper.GetString("database.username")
	databasePassword := viper.GetString("database.password")
	databaseName := viper.GetString("database.dbname")

	fmt.Println("Server Host:", serverHost)
	fmt.Println("Server Port:", serverPort)
	fmt.Println("Database Username:", databaseUsername)
	fmt.Println("Database Password:", databasePassword)
	fmt.Println("Database Name:", databaseName)
}

func WriteValuesToConfigFile() {

	viper.AddConfigPath(".")
	viper.SetConfigName("config2")
	viper.SetConfigType("yaml")

	viper.Set("server.host", "localhost")
	viper.Set("server.port", 8080)
	viper.Set("database.username", "admin")
	viper.Set("database.password", "password")
	viper.Set("database.dbname", "mydb")

	if err := viper.WriteConfigAs("config2.yaml"); err != nil {
		fmt.Println("Error writing config file,", err)
	}
}
