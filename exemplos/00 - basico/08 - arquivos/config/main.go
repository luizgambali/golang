package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func main() {

	var cfg Config
	readFile(&cfg)
	//readEnv(&cfg)

	fmt.Printf("Server.Host: %v\n", cfg.Server.Host)
	fmt.Printf("Server.Port: %v\n", cfg.Server.Port)

	fmt.Printf("Database.Username: %v\n", cfg.Database.Username)
	fmt.Printf("Database.Password: %v\n", cfg.Database.Password)

}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
