package config

import "sync"

type (
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
)

var (
	instance *Config
	once     sync.Once
)
