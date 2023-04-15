package scheduler

import (
	"context"
	"time"
)

func (s *scheduler) TemporalArbitrage(ctx context.Context, delay time.Duration) {
	ticker := time.NewTicker(delay)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := s.bundle.Clear(ctx)
				if err != nil {
					s.log.Errorf("Scheduler: error clear bundle table: %v", err)

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
					s.log.Errorf("Scheduler: error save bundle - %v : %v", bundle, err)

					break
				}
			}
		}
	}()
}
