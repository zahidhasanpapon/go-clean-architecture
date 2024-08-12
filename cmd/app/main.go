package main

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-go-arch/config"
)

func main() {
	logger := setUpLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.String("error", err.Error()))
		}
	}(logger)

	// Initialize the configuration using Viper
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal("Error initializing config", zap.String("error", err.Error()))
	}

	// Optionally override configuration with etcd
	if err := config.LoadConfigFromEtcd(cfg); err != nil {
		logger.Info("Could not load config from etcd", zap.String("error", err.Error()))
	}

	// Set up signal handling for configuration reload
	reloadChan := make(chan os.Signal, 1)
	signal.Notify(reloadChan, syscall.SIGHUP)

	go func() {
		for {
			<-reloadChan
			logger.Info("Reloading configuration...")
			if err := config.LoadConfigFromEtcd(cfg); err != nil {
				logger.Info("Could not reload config from etcd", zap.String("error", err.Error()))
			}
		}
	}()

	// Correlation IDs for Traceability
	/*	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logger.With(
			zap.String("request_id", getCorrelationID(ctx)),
			zap.String("path", r.URL.Path),
		)

		logger.Info("Handling request")
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			return
		}
	})*/

	// Start service
	logger.Info("Starting service", zap.String("app_name", cfg.AppName), zap.Int("port", cfg.Port))
}

func setUpLogger() *zap.Logger {
	logLevel := zapcore.InfoLevel
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		if err := logLevel.Set(envLogLevel); err != nil {
			logLevel = zapcore.InfoLevel
		}
	}

	loggerConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: false,
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, _ := loggerConfig.Build()
	return logger
}

func getCorrelationID(ctx context.Context) string {
	if correlationID, ok := ctx.Value("correlation_id").(string); ok {
		return correlationID
	}
	return "unknown"
}

func correlationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), "correlation_id", correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
