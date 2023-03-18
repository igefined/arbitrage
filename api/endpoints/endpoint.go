package endpoints

import (
	"net/http"

	"github.com/igilgyrg/arbitrage/config"
	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use"
)

type Endpoint interface {
	Bundles() http.HandlerFunc
	Status() http.HandlerFunc
}

type endpoint struct {
	cfg    *config.Config
	logger *log.Logger
	use    use.UseCase
}

func New(cfg *config.Config, logger *log.Logger, use use.UseCase) Endpoint {
	return &endpoint{cfg: cfg, logger: logger, use: use}
}
