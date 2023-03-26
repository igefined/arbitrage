package inspector

import (
	"context"
	"fmt"
	"math"
)

type spread struct {
	ExchangeName string
	Price        float64
}

func (s *service) Inspect(ctx context.Context) {
	s.log.Infof("running inspection")
	symbols, err := s.Symbols(ctx)
	if err != nil {
		s.log.Errorf("Inspector: error of getting symbols: %v", err)

		return
	}

	if len(s.exchangers) < 1 {
		return
	}

	spreads := make(map[string][]spread, len(symbols))

	for _, symb := range symbols {
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
			for _, s := range v {
				if tmp.Price != s.Price {
					percent := (s.Price - tmp.Price) / tmp.Price * 100
					if math.Abs(percent) > 1 {
						fmt.Printf("%s: %s - %f\n", k, s.ExchangeName, s.Price)
					}
				}
			}
		}
	}

	s.log.Infof("inspection stopped")
}
