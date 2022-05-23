package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App  `yaml:"APP"`
		HTTP `yaml:"HTTP"`
		Log  `yaml:"LOGGER"`
		PG   `yaml:"POSTGRES"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"NAME"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"VERSION" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true"  yaml:"HTTP_PORT"  env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true"  yaml:"LOG_LEVEL"  env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true"  yaml:"PG_POOL_MAX"  env:"PG_POOL_MAX"`
		URL     string `env-required:"true"  yaml:"PG_URL"       env:"PG_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigFile("./config/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
