package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Config

type (
	Config struct {
		App      `mapstructure:"APP"`
		HTTP     `mapstructure:"HTTP"`
		Log      `mapstructure:"LOGGER"`
		Postgres `mapstructure:"POSTGRES"`
	}

	// App -.
	App struct {
		Name    string `mapstructure:"NAME"`
		Version string `mapstructure:"VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `mapstructure:"PORT"`
	}

	// Log -.
	Log struct {
		Level string `mapstructure:"LEVEL"`
	}

	// Postgres -.
	Postgres struct {
		PoolMax     int    `mapstructure:"POOL_MAX"`
		DatabaseUrl string `mapstructure:"DATABASE_URL"`
	}
)

func New() (*Config, error) {
	viper.SetConfigFile("../../config/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return Cfg, nil
}
