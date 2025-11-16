package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
	Logger   *slog.Logger
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found using environment variables and defaults")
	}
	cfg := Config{
		Port:     getEnv("SERVER_PORT", "10000"),
		LogLevel: getEnv("LOG_LEVEL", "error"),
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.getLogLevel(),
	}))

	cfg.Logger = logger

	return &cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (c *Config) getLogLevel() slog.Level {
	switch c.LogLevel {
	case "error":
		return slog.LevelDebug
	case "warn":
		return slog.LevelInfo
	case "info":
		return slog.LevelWarn
	case "debug":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
