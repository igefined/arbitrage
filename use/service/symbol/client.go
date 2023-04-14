package symbol

import (
	"context"
	"sync"

	"github.com/igilgyrg/arbitrage/log"
	"github.com/igilgyrg/arbitrage/use/integration/ninja"
)

type Service interface {
	Symbols(ctx context.Context) []string
	Upgrade(ctx context.Context)
}

type client struct {
	log   *log.Logger
	ninja ninja.Client

	mutex   sync.RWMutex
	symbols []string
}

func New(log *log.Logger, ninja ninja.Client) (Service, error) {
	symbols, err := ninja.CryptoSymbols(context.Background())
	if err != nil {
		return nil, err
	}

	tempSymbols := make([]string, 0, len(symbols))
	for i := range symbols {
		symbol := symbols[i]
		if validate(symbol) {
			tempSymbols = append(tempSymbols, symbol)
		}
	}

	return &client{
		log:     log,
		ninja:   ninja,
		symbols: tempSymbols,
	}, nil
}
