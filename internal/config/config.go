package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`
}

type Config struct {
	Port           string `env:"PORT" envDefault:"8080"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`
	DB             DBConfig
	AgifyURL       string `env:"AGIFY_URL" envDefault:"https://api.agify.io"`
	GenderizeURL   string `env:"GENDERIZE_URL" envDefault:"https://api.genderize.io"`
	NationalizeURL string `env:"NATIONALIZE_URL" envDefault:"https://api.nationalize.io"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
