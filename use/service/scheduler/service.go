package scheduler

import (
	"context"
	"time"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/symbol"
)

type Service interface {
	TemporalArbitrage(ctx context.Context, delay time.Duration)
	TemporalSymbols(ctx context.Context, delay time.Duration)
}

type scheduler struct {
	log *log.Logger

	inspector inspector.Service
	bundle    bundle.Service
	symbols   symbol.Service
}

func New(log *log.Logger, inspector inspector.Service, bundle bundle.Service, symbols symbol.Service) Service {
	return &scheduler{log: log, inspector: inspector, bundle: bundle}
}
