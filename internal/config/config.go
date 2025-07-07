package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	DBName            string `yaml:"dbname"`
	UserTableName     string `yaml:"users_table"`
	TaskListTableName string `yaml:"tasks_lists_table"`
	TaskTableName     string `yaml:"tasks_table"`
	SSLMode           string `yaml:"sslmode"`
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

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
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
