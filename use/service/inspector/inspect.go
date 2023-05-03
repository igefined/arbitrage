package inspector

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/utils/usymbol"
)

const percentageDifference = 2

func (s *service) Inspect(ctx context.Context) {
	s.log.Infof("running inspection")

	if len(s.exchangers) < 1 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		activeSymbols := s.symbols.Symbols(ctx)

		for _, symb := range activeSymbols {
			spreadsInfo := make([]domain.SpreadInfo, 0, len(s.exchangers))

			for _, e := range s.exchangers {
				ticker, tickerErr := e.DailyTicker(ctx, symb)
				if tickerErr != nil {
					continue
				}

				cryptoName, _ := usymbol.Split(symb)
				depositNetworks := e.DepositNetwork(ctx, cryptoName)
				withdrawNetworks := e.WithdrawNetwork(ctx, cryptoName)

				spreadsInfo = append(spreadsInfo, domain.SpreadInfo{
					ExchangeName:     e.Name(),
					Price:            ticker.Price,
					DepositNetworks:  depositNetworks,
					WithdrawNetworks: withdrawNetworks,
				})
			}

			spread := domain.Spreads{
				Symbol:  symb,
				Spreads: spreadsInfo,
			}

			s.spreads <- spread
		}

		wg.Done()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case spread := <-s.spreads:
				symbol := spread.Symbol
				spreads := spread.Spreads

				for i := range spreads {
					tmp := spreads[i]
					for _, spr := range spreads {
						if tmp.Price != spr.Price {
							percent := (spr.Price - tmp.Price) / tmp.Price * 100
							percentAsString := fmt.Sprintf("%.3f", percent)
							percentFloat, err := strconv.ParseFloat(percentAsString, 64)
							if err != nil {
								s.log.Error(err)

								continue
							}

							if percent > percentageDifference {
								depositNetworks := tmp.DepositNetworks
								withdrawNetworks := spr.WithdrawNetworks

								for _, depN := range depositNetworks {
									for _, withN := range withdrawNetworks {
										if strings.EqualFold(depN, withN) {
											bundle := domain.Bundle{
												Symbol:               symbol,
												ExchangeFrom:         tmp.ExchangeName,
												PriceFrom:            tmp.Price,
												ExchangeTo:           spr.ExchangeName,
												PriceTo:              spr.Price,
												PercentageDifference: percentFloat,
												WithdrawNetworks:     spr.WithdrawNetworks,
												DepositNetworks:      spr.DepositNetworks,
											}

											s.bundles <- bundle
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}()

	wg.Wait()
	s.log.Infof("inspection stopped")
}
