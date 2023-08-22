package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/kiramishima/shining_guardian/src/domain"
)

func Load() (*domain.Configuration, error) {
	var cfg domain.Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// NewConfig creates and load config
func NewConfig() *domain.Configuration {
	cfg, err := Load()
	if err != nil {
		log.Printf("Can't load the configuration. Error: %s", err.Error())
	}

	return cfg
}
