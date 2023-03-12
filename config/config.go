package config

import (
	"context"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Config struct {
	Name string `config:"APP_NAME"`
	Port uint   `config:"PORT"`
}

func New() *Config {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cfg := &Config{}
	loader := confita.NewLoader(env.NewBackend())
	err := loader.Load(ctx, cfg)
	if err != nil {
		panic("config has not loaded")
	}

	return cfg
}
