package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

// List is a custom type to be used for parsing yaml lists
type List []string

// SetValue is required for this custom setter; maps a comma separated string into a slice of strings
func (f *List) SetValue(s string) error {
	*f = strings.Split(s, ", ")
	return nil
}

// APIConfig holds the configuration of the api
type APIConfig struct {
	Port        string `yaml:"port" env:"GATEWAY_API_PORT" env-default:"5009"`
	CorsOrigins List   `yaml:"allowed_origins" env:"GATEWAY_API_ALLOWED_ORIGINS" env-default:""`
}

// DatabaseConfig is a representation of the database configuration
type DatabaseConfig struct {
	User     string `yaml:"user" env:"GATEWAY_DB_USER"`
	Password string `yaml:"password" env:"GATEWAY_DB_PASSWORD"`
	DB       string `yaml:"db" env:"GATEWAY_DB_DB_NAME" env-default:"gateway"`
}

// PicoConfig represents the configuration for ganesh
type PicoConfig struct {
	LogLevel string         `yaml:"log_level" env:"GATEWAY_LOG_LEVEL" env-default:"info"`
	API      APIConfig      `yaml:"api"`
	Database DatabaseConfig `yaml:"database"`
}

// GetConfig retrieves the configuration for the project
func GetConfig() PicoConfig {
	var config PicoConfig
	configPath := os.Getenv("PICO_CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		fmt.Println("WARNING: Could not read config file; Continuing with defaults")

		// panic(fmt.Sprintf("Could not read config file: %s", err))
	}
	return config
}
