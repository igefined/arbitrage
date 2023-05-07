package okx

import (
	"context"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "okx"

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		cfg    *config
		logger *log.Logger
	}

	config struct {
		ApiKey     string `config:"OKX_API_KEY"`
		SecretKey  string `config:"OKX_SECRET_KEY"`
		Passphrase string `config:"OKX_PASSPHRASE"`
	}
)

func New(logger *log.Logger) exchangers.Client {
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
		"https://www.okx.com",
		"https://aws.okx.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, cfg: cfg, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
