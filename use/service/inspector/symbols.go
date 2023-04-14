package inspector

import (
	"context"
	"sync"
)

type symbols struct {
	symbols []string
	mutex   sync.RWMutex
}

func (s *service) Symbols(ctx context.Context) []string {
	return s.symbols.Symbols(ctx)
}
