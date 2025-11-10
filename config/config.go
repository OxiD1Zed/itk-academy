package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Postgres  PostgresConfig
	AppConfig AppConfig
}

type PostgresConfig struct {
	Port           uint16
	MaxConnections int
	AcquireTimeout time.Duration
	Host           string
	Username       string
	Password       string
	Database       string
}

type AppConfig struct {
	Port uint16
	Host string
}

func NewConfig() *Config {
	return &Config{
		PostgresConfig{
			Port:           uint16(getEnvAsInt("POSTGRES_PORT", 5432)),
			MaxConnections: getEnvAsInt("POSTGRES_MAX_CONNECTIONS", 10),
			AcquireTimeout: time.Duration(getEnvAsInt("POSTGRES_ACQUIRE_TIMEOUT", 300)) * time.Millisecond,
			Host:           getEnv("POSTGRES_HOST", "localhost"),
			Username:       getEnv("POSTGRES_USERNAME", "postgres"),
			Password:       getEnv("POSTGRES_PASSWORD", "postgres"),
			Database:       getEnv("POSTGRES_DATABASE", "database"),
		},
		AppConfig{
			Port: uint16(getEnvAsInt("APP_PORT", 8080)),
			Host: getEnv("APP_HOST", "0.0.0.0"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if valStr, exists := os.LookupEnv(key); exists {
		val, err := strconv.ParseInt(valStr, 10, 64)
		if err != nil {
			return defaultVal
		}

		return int(val)
	}

	return defaultVal
}
