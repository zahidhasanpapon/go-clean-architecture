package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func InitConfig() (*Config, error) {
	var cfg Config

	// Set default values
	viper.SetDefault("app_name", "Go Microservice")
	viper.SetDefault("port", 8080)

	// Load environment variables
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	// Load configuration from file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Override with environment variables
	if port, exists := os.LookupEnv("APP_PORT"); exists {
		if portInt, err := strconv.Atoi(port); err == nil {
			cfg.Port = portInt
		} else {
			log.Printf("Invalid port value: %s", port)
		}
	}

	return &cfg, nil
}
