package bybit

import (
	"context"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const (
	ExchangeName = "bybit"
	recvWindow   = 5_000
)

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		cfg    *config
		logger *logger.Logger
	}

	config struct {
		ApiKey    string `config:"BYBIT_API_KEY"`
		SecretKey string `config:"BYBIT_SECRET_KEY"`
	}
)

func New(logger *logger.Logger) exchangers.Client {
	cfg := &config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, cfg); err != nil {
		logger.Error(err)
	}

	httpClient := &http.Client{
		Timeout: exchangers.ProvTimeoutSec * time.Second,
	}

	hosts := []string{
		"https://api.bybit.com",
		"https://api.bytick.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, cfg: cfg, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
