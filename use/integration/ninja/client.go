package ninja

import (
	"context"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/arbitrage/log"
)

const defaultEndpoint = "https://api.api-ninjas.com"

type Client interface {
	CryptoSymbols(ctx context.Context) ([]string, error)
}

type (
	config struct {
		Endpoint string `config:"NINJA_HOST"`
		ApiKey   string `config:"NINJA_API_KEY"`
	}

	client struct {
		log *log.Logger
		cfg *config

		httpClient *http.Client
	}
)

func New(log *log.Logger) (Client, error) {
	cfg := &config{
		Endpoint: defaultEndpoint,
		// ApiKey
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := confita.NewLoader(env.NewBackend()).Load(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &client{
		log:        log,
		cfg:        cfg,
		httpClient: &http.Client{},
	}, nil
}
