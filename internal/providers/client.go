package providers

import "context"

type Client interface {
	GetBySymbol(ctx context.Context)
}
