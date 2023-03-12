package endpoints

import (
	"net/http"

	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"
)

type Endpoint interface {
	Bundles() http.HandlerFunc
	Status() http.HandlerFunc
}

type endpoint struct {
	cfg    *config.Config
	logger *log.Logger
}

func New(cfg *config.Config, logger *log.Logger) Endpoint {
	return &endpoint{cfg: cfg, logger: logger}
}
