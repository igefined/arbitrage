package inspector

import (
	"context"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/binance"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/bybit"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/huobi"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/kucoin"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/mexc"
	"github.com/igilgyrg/arbitrage/use/service/symbol"
)

type Service interface {
	Inspect(ctx context.Context)
	Bundles() chan domain.Bundle
}

type service struct {
	log     *log.Logger
	symbols symbol.Service

	exchangers []exchangers.Client
	bundles    chan domain.Bundle
	spreads    chan domain.Spreads
}

func New(log *log.Logger, symbols symbol.Service, exchangers []exchangers.Client) Service {
	return &service{log: log, exchangers: exchangers, symbols: symbols, bundles: make(chan domain.Bundle, 1), spreads: make(chan domain.Spreads, 1)}
}

func (s *service) Bundles() chan domain.Bundle {
	return s.bundles
}

func DefaultExchangers(log *log.Logger) []exchangers.Client {
	return []exchangers.Client{binance.New(log), bybit.New(log), mexc.New(log), huobi.New(log), kucoin.New(log)}
}
