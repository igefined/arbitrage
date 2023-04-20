package exchangers

import (
	"context"
	"errors"

	"github.com/igilgyrg/arbitrage/use/domain"
)

const ProvTimeoutSec = 5

var (
	ErrUnavailable    = errors.New("server not available")
	ErrSymbolNotFound = errors.New("symbol not found")
)

type Client interface {
	DailyTicker(ctx context.Context, symbol string) (*domain.DailyTicker, error)
	IsWithdraw(ctx context.Context, symbol string) bool
	IsDeposit(ctx context.Context, symbol string) bool

	Name() string
}
