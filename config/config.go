package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type (
	// Config holds the application configuration
	Config struct {
		AppName     string     `mapstructure:"app_name"`
		Port        int        `mapstructure:"port"`
		LogLevel    string     `mapstructure:"log_level"`
		DatabaseURL string     `mapstructure:"database_url"`
		RedisURL    string     `mapstructure:"redis_url"`
		EtcdConfig  EtcdConfig `mapstructure:"etcd"`
	}

	EtcdConfig struct {
		Endpoints []string `mapstructure:"endpoints"`
		Username  string   `mapstructure:"username"`
		Password  string   `mapstructure:"password"`
	}

	DatabaseConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
	// Add more configuration struct here
)

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton instance of the configuration
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			AppName:     "go-microservice",
			Port:        8080,
			LogLevel:    "debug",
			DatabaseURL: "localhost:5432",
			RedisURL:    "localhost:6379",
			EtcdConfig: EtcdConfig{
				Endpoints: []string{"localhost:2379"},
				Username:  "",
				Password:  "",
			},
		}
		InitConfig()
	})
	return instance
}

func initConfig() {
	// Initialize local configuration
	initLocalConfig()

	// Initialize etcd configuration
	initEtcdConfig()

	// Initialize Vault configuration
	initVaultConfig()

	// Unmarshal the configuration
	if err := viper.Unmarshal(instance); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// Set up configuration reload for the local changes
	watchLocalConfigFile()
}
