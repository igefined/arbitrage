package binance

import (
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "binance"

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		logger *log.Logger
	}
)

func New(logger *log.Logger) exchangers.Client {
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

	return &client{httpClient: httpClient, hosts: hosts, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
