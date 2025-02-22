package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds application configurations.
type Config struct {
	AppPort    int    `yaml:"app_port"`
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DebugMode  bool   `yaml:"debug_mode"`
	// Add more configuration fields as necessary
}

func LoadConfig() (*Config, error) {
	env := getEnv("APP_ENV", "dev")
	configFile := fmt.Sprintf("./config/environments/%s.yml", env)

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %w", err)
	}

	// Optionally validate the config or set defaults here

	return &cfg, nil
}

func getEnv(key, fallback string) string {
	value := fallback
	if val, exists := os.LookupEnv(key); exists {
		value = val
	}
	return value
}
