package binance

import (
	"context"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/binance/response"
)

const (
	ExchangeName = "binance"
	recvWindow   = 50_000
)

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		cfg          *config
		logger       *log.Logger
		allCoinsInfo map[string]response.CoinInformation
	}

	config struct {
		ApiKey    string `config:"BINANCE_API_KEY"`
		SecretKey string `config:"BINANCE_SECRET_KEY"`
	}
)

func New(logger *log.Logger) exchangers.Client {
	cfg := &config{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := confita.NewLoader(env.NewBackend()).Load(ctx, cfg)
	if err != nil {
		logger.Error(err)
	}

	httpClient := &http.Client{
		Timeout: exchangers.ProvTimeoutSec * time.Second,
	}

	hosts := []string{
		"https://api.binance.com",
		"https://api1.binance.com",
		"https://api2.binance.com",
		"https://api3.binance.com",
		"https://api4.binance.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, cfg: cfg, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
