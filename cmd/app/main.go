package main

import (
	"os"
	"os/signal"
	"syscall"
	"test-go-arch/config"
	"test-go-arch/pkg/logger"
)

func main() {
	// Initialize the logger
	log := logger.New(logger.InfoLevel)

	// Initialize the configuration using Viper
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error initializing config", "error", err.Error())
	}

	// Optionally override configuration with etcd
	if err := config.LoadConfigFromEtcd(cfg); err != nil {
		log.Info("Could not load config from etcd", "error", err.Error())
	}

	// Set up signal handling for configuration reload
	reloadChan := make(chan os.Signal, 1)
	signal.Notify(reloadChan, syscall.SIGHUP)

	go func() {
		for {
			<-reloadChan
			log.Info("Reloading configuration...")
			if err := config.LoadConfigFromEtcd(cfg); err != nil {
				log.Info("Could not reload config from etcd", "error", err.Error())
			}
		}
	}()

	// Start service
	log.Info("Starting service", "app_name", cfg.AppName, "port", cfg.Port)
}
