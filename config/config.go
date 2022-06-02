package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *Config

type (
	Config struct {
		App      `json:"app" mapstructure:"app"`
		HTTP     `json:"http" mapstructure:"http"`
		Logger   `json:"logger" mapstructure:"logger"`
		Postgres `json:"postgres" mapstructure:"postgres"`
	}

	// App -.
	App struct {
		Name    string `json:"name" mapstructure:"name"`
		Version string `json:"version" mapstructure:"version"`
	}

	Logger struct {
		Level string `json:"level" mapstructure:"level"`
	}

	// HTTP -.
	HTTP struct {
		Port string `json:"port" mapstructure:"port"`
	}

	// Postgres -.
	Postgres struct {
		PoolMax     int    `json:"pool_max" mapstructure:"pool_max"`
		DatabaseUrl string `json:"database_url" mapstructure:"database_url"`
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
