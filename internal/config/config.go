package config

import (
	"fmt"
	"github.com/cristalhq/aconfig"
	"strings"
)

type Config struct {
	ListenAddress string `default:"127.0.0.1:5000" env:"LISTEN_ADDRESS"`
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) Load() error {
	err := aconfig.LoaderFor(cfg, aconfig.Config{
		SkipFiles: true,
		SkipFlags: true,
		EnvPrefix: "WEB",
	}).Load()
	if err != nil {
		return err
	}

	return cfg.Validate()
}

func (cfg *Config) Validate() error {
	if strings.TrimSpace(cfg.ListenAddress) == "" {
		return fmt.Errorf("ListenAddress must be defined")
	}

	return nil
}
