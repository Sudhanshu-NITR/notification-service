package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.App.Name = getEnv("APP_NAME", "notification-service")
	cfg.App.Env = getEnv("APP_ENV", "development")
	cfg.App.Port = getEnv("APP_PORT", "8080")

	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")
	cfg.Database.Name = getEnv("DB_NAME", "notifications")

	cfg.Queue.URL = getEnv("QUEUE_URL", "")

	cfg.Security.APIKeySecret = getEnv("API_KEY_SECRET", "")

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
