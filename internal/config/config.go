package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type AppConfig struct {
	Env       string        `yaml:"env"`
	Port      string        `yaml:"port"`
	JWTSecret string        `yaml:"jwt_secret"`
	TokenTTL  time.Duration `yaml:"token_ttl"`
}

type Config struct {
	App            AppConfig
	PostgresConfig PostgresConfig
}

func (cfg *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@localhost:%v/%s", cfg.User, cfg.Password, cfg.Port, cfg.DBName,
	)
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %s", err)
	}
	cfg := Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}
	return &cfg, nil
}
