package config

type Config struct {
	Server struct {
		GRPCPort    string `env:"GRPC_PORT"`
		HTTPPort    string `env:"HTTP_PORT"`
		MetricsPort string `env:"METRICS_PORT"`
	}
	MySQL MySQL
}

// MySQL is
type MySQL struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}
