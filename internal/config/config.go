package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		// GRPCPort    string `env:"GRPC_PORT"`
		HTTPPort string `env:"HTTP_PORT"`
		// MetricsPort string `env:"METRICS_PORT"`
	}
	MySQL MySQL `env:""`
}

// MySQL is
type MySQL struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

// LoadConfig is
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("[internal][config][LoadConfig]: %w", err)
	}

	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("[internal][config][Process]: %w", err)
	}

	// Check http Port
	httpPort, exists := os.LookupEnv("HTTP_PORT")
	if !exists {
		return nil, fmt.Errorf("HTTP_PORT NOT EXIST: %w", err)
	}
	c.Server.HTTPPort = httpPort

	// Validate the config
	err = validator.New().Struct(c)
	if err != nil {
		return nil, fmt.Errorf("[internal][config][validate]: %w", err)
	}

	return &c, nil
}
