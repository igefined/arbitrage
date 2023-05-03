package huobi

import (
	"context"
	"fmt"
	"strings"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	resp, err := c.market.GetLatestTrade(strings.ToLower(symbol))
	if err != nil {
		err = fmt.Errorf("huobi daily ticker request: %v", err)

		return
	}

	if resp.Data == nil {
		err = fmt.Errorf("huobi daily ticker: nil result")

		return
	}

	if len(resp.Data) == 0 {
		err = fmt.Errorf("huobi daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	firstTicker := resp.Data[0]
	price, _ := firstTicker.Price.BigFloat().Float64()
	ticker = &domain.DailyTicker{
		Symbol: strings.ToUpper(symbol),
		Price:  price,
	}

	return
}
