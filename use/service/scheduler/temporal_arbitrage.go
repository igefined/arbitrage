package scheduler

import (
	"context"
	"time"
)

func (s *scheduler) TemporalArbitrage(ctx context.Context, delay time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.Tick(delay):
				err := s.bundle.Clear(ctx)
				if err != nil {
					s.log.Error("Scheduler: error clear bundle table: %v", err)

					break
				}

				s.inspector.Inspect(ctx)
			}
		}
	}()

	go func() {
		bundles := s.inspector.Bundles()

		for {
			select {
			case <-ctx.Done():
				s.log.Error("Scheduler: context done")

				return
			case bundle := <-bundles:
				err := s.bundle.Save(ctx, &bundle)
				if err != nil {
					s.log.Error("Scheduler: error save bundle - %v : %v", bundle, err)

					break
				}
			}
		}
	}()
}
