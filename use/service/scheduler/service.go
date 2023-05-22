package scheduler

import (
	"context"
	"time"

	"github.com/igdotog/core/logger"
	"github.com/igilgyrg/arbitrage/use/integration/bot"
	"github.com/igilgyrg/arbitrage/use/service/bundle"
	"github.com/igilgyrg/arbitrage/use/service/inspector"
	"github.com/igilgyrg/arbitrage/use/service/symbol"
)

type Service interface {
	TemporalArbitrage(ctx context.Context, delay time.Duration)
	TemporalSymbols(ctx context.Context, delay time.Duration)
}

type scheduler struct {
	log *logger.Logger

	inspector inspector.Service
	bundle    bundle.Service
	symbols   symbol.Service
	bot       bot.Client
}

func New(log *logger.Logger, inspector inspector.Service, bundle bundle.Service, symbols symbol.Service, bot bot.Client) Service {
	return &scheduler{log: log, inspector: inspector, bundle: bundle, symbols: symbols, bot: bot}
}
