package providers

import (
	"context"
	"errors"

	"github.com/igilgyrg/arbitrage/internal/domain"
)

const (
	ProvTimeoutSec   = 5
	SymbolStableCoin = "USDT"
)

var (
	ErrUnavailable    = errors.New("server not available")
	ErrSymbolNotFound = errors.New("symbol not found")
)

type Client interface {
	DailyTicker(ctx context.Context, symbol string) (*domain.DailyTicker, error)
}
