package config

import (
	"context"
	"log"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Config struct {
	Name  string `config:"APP_NAME"`
	Port  uint   `config:"PORT"`
	DBUrl string `config:"DB_URL,required"`
}

func New() *Config {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cfg := &Config{}
	loader := confita.NewLoader(env.NewBackend())
	err := loader.Load(ctx, cfg)
	if err != nil {
		log.Fatal("config has not loaded")
	}

	return cfg
}
