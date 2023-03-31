package huobi

import (
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

const ExchangeName = "huobi"

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
		"https://api.huobi.pro",
		"https://api-aws.huobi.pro",
	}

	return &client{httpClient: httpClient, hosts: hosts, logger: logger}
}

func (c *client) Name() string {
	return ExchangeName
}
