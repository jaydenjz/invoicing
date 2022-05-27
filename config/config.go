package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Config

type (
	Config struct {
		App      `mapstructure:"app"`
		HTTP     `mapstructure:"http"`
		Log      `mapstructure:"logger"`
		Postgres `mapstructure:"postgres"`
	}

	// App -.
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
	}

	// HTTP -.
	HTTP struct {
		Port string `mapstructure:"port"`
	}

	// Log -.
	Log struct {
		Level string `mapstructure:"level"`
	}

	// Postgres -.
	Postgres struct {
		PoolMax     int    `mapstructure:"pool_max"`
		DatabaseUrl string `mapstructure:"database_url"`
	}
)

func New() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config/")
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
