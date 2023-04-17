package scheduler

import (
	"context"
	"time"

	"github.com/igilgyrg/arbitrage/use/integration/bot"
)

const chatId = 5287037408

func (s *scheduler) TemporalArbitrage(ctx context.Context, delay time.Duration) {
	ticker := time.NewTicker(delay)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
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

				botMsg := bot.Message{
					ChatId:  chatId,
					Content: bundle.String(),
				}
				err = s.bot.Send(ctx, botMsg)
				if err != nil {
					s.log.Errorf("Scheduler: error send bundle to bot - %v : %v", bundle, err)

					break
				}
			}
		}
	}()
}
