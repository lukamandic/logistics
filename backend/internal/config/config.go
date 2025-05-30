package config

import (
	"fmt"
	"os"
)

type Config struct {
	UIURL      string
	TableName  string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		UIURL:      getEnvOrDefault("UI_URL", "http://localhost:5173"),
		TableName:  getEnvOrDefault("TABLE_NAME", "logistics"),
		ServerPort: getEnvOrDefault("SERVER_PORT", "8080"),
	}

	if config.TableName == "" {
		return nil, fmt.Errorf("TABLE_NAME environment variable is required")
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}