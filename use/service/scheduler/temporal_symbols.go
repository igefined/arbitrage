package scheduler

import (
	"context"
	"time"
)

func (s *scheduler) TemporalSymbols(ctx context.Context, delay time.Duration) {
	ticker := time.NewTicker(delay)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.symbols.Upgrade(ctx)
			}
		}
	}()
}
