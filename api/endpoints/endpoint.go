package endpoints

import (
	"net/http"

	"github.com/igdotog/core/config"
	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use"
)

type Endpoint interface {
	Bundles() http.HandlerFunc
	Status() http.HandlerFunc
}

type endpoint struct {
	cfg    *config.BaseConfig
	logger *logger.Logger
	use    use.UseCase
}

func New(cfg *config.BaseConfig, logger *logger.Logger, use use.UseCase) Endpoint {
	return &endpoint{cfg: cfg, logger: logger, use: use}
}
