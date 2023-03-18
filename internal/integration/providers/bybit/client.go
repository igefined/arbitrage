package bybit

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
		"https://api.bybit.com",
		"https://api.bytick.com",
	}

	return &client{httpClient: httpClient, hosts: hosts, logger: logger}
}
