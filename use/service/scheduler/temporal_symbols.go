package scheduler

import (
	"context"
	"time"
)

func (s *scheduler) TemporalSymbols(ctx context.Context, delay time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.Tick(delay):
				s.symbols.Upgrade(ctx)
			}
		}
	}()
}
