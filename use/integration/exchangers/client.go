package exchangers

import (
	"context"
	"errors"

	"github.com/igilgyrg/arbitrage/use/domain"
)

const ProvTimeoutSec = 5

var (
	ErrUnavailable         = errors.New("server not available")
	ErrSymbolNotFound      = errors.New("symbol not found")
	ErrEmptyNetworks       = errors.New("empty networks")
	ErrInvalidApiSignature = errors.New("invalid api signature")
)

type Client interface {
	DailyTicker(ctx context.Context, symbol string) (*domain.DailyTicker, error)
	WithdrawNetwork(ctx context.Context, symbol string) []string
	DepositNetwork(ctx context.Context, symbol string) []string

	Name() string
}
