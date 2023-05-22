package kucoin

import (
	"net/http"
	"time"

	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "kucoin"

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		logger *logger.Logger
	}
)

func New(logger *logger.Logger) exchangers.Client {
	httpClient := &http.Client{
		Timeout: exchangers.ProvTimeoutSec * time.Second,
	}

	hosts := []string{
		"https://api.kucoin.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
