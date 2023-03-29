package inspector

import (
	"context"

	"github.com/igilgyrg/arbitrage/use/domain"
)

const percentageDifference = 0

type spread struct {
	ExchangeName string
	Price        float64
}

func (s *service) Inspect(ctx context.Context) {
	s.log.Infof("running inspection")

	activeSymbols, err := s.Symbols()
	if err != nil {
		s.log.Errorf("Inspector: error of getting symbols: %v", err)

		return
	}

	if len(s.exchangers) < 1 {
		return
	}

	spreads := make(map[string][]spread, len(activeSymbols))

	for _, symb := range activeSymbols {
		sprs := make([]spread, 0, len(s.exchangers))

		for _, e := range s.exchangers {
			ticker, tickerErr := e.DailyTicker(ctx, symb)
			if tickerErr != nil {
				continue
			}

			sprs = append(sprs, spread{
				ExchangeName: e.Name(),
				Price:        ticker.Price,
			})
		}

		spreads[symb] = sprs
	}

	for k, v := range spreads {
		for i := 0; i < len(v); i++ {
			tmp := v[i]
			for _, spr := range v {
				if tmp.Price != spr.Price {
					percent := (spr.Price - tmp.Price) / tmp.Price * 100
					if percent > percentageDifference {
						bundle := domain.Bundle{
							Symbol:               k,
							ExchangeFrom:         tmp.ExchangeName,
							ExchangeTo:           spr.ExchangeName,
							PercentageDifference: percent,
						}

						s.bundles <- bundle
					}
				}
			}
		}
	}

	s.log.Infof("inspection stopped")
}
