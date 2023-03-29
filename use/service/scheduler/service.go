package scheduler

import (
	"context"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
)

type Service interface {
	TemporalArbitrage(ctx context.Context, delay time.Duration)
}

type scheduler struct {
	log *log.Logger

	inspector inspector.Service
	bundle    bundle.Service
}

func New(log *log.Logger, inspector inspector.Service, bundle bundle.Service) Service {
	return &scheduler{log: log, inspector: inspector, bundle: bundle}
}
