package bybit

import (
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "bybit"

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
		"https://api.bybit.com",
		"https://api.bytick.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
