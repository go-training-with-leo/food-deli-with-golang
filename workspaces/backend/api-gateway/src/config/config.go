package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	App Config
)

type Config struct {
	DatabaseType     string `mapstructure:"DB_TYPE"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     int    `mapstructure:"DB_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`

	Port string `mapstructure:"PORT"`
}

func init() {
	LoadedConfAppFromTheEnv, err := loadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config.\n", err)
	}

	App = LoadedConfAppFromTheEnv

	fmt.Println("Loaded config!")
}

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Cannot load .env file.\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Unable to decode into struct.\n", err)
	}

	fmt.Println("Loaded env vars from .env file!")

	return
}
