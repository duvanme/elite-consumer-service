package config

import (
	"os"
)

type Config struct {
	DatabaseUrl  string
	Port         string
	KafkaBrokers []string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseUrl:  os.Getenv("DATABASE_URL"),
		Port:         getEnv("PORT", "8080"),
		KafkaBrokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
