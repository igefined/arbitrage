package binance

import (
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/internal/integration/providers"
	"github.com/igilgyrg/arbitrage/log"
)

type (
	client struct {
		httpClient *http.Client
		hosts      []string

		logger *log.Logger
	}
)

func New(logger *log.Logger) providers.Client {
	httpClient := &http.Client{
		Timeout: providers.ProvTimeoutSec * time.Second,
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
