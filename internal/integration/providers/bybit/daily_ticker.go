package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/internal/domain"
	"github.com/igilgyrg/arbitrage/internal/integration/providers"
)

func (c client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	symbol = fmt.Sprintf("%s%s", symbol, providers.SymbolStableCoin)
	query := fmt.Sprintf("%s?symbol=%s&category=spot", "v5/market/tickers", symbol)
	resp, err := providers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, nil)
	if err != nil {
		err = fmt.Errorf("bybit daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("bybit daily ticker response status: %s", resp.Status)

		return
	}

	response := Response{}
	response.Result = &TickerResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		err = fmt.Errorf("bybit daily ticker decoder: %v", err)

		return
	}

	if response.Code != 0 {
		switch response.Code {
		case 10001:
			err = fmt.Errorf("bybit daily ticker: %w", providers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("bybit daily ticker error response: %s", response.Message)
		}

		return
	}

	if response.Result == nil {
		err = fmt.Errorf("bybit daily ticker: nil result")

		return
	}

	tickerResponse, ok := response.Result.(*TickerResponse)
	if !ok {
		err = fmt.Errorf("bybit daily ticker decoder: cannot json decode result")

		return
	}

	if tickerResponse.Category != "spot" {
		err = fmt.Errorf("bybit daily ticker: not spot account")

		return
	}

	if len(tickerResponse.List) == 0 {
		err = fmt.Errorf("bybit daily ticker: %w", providers.ErrSymbolNotFound)

		return
	}

	ticker = tickerResponse.List[0].ToResponse()

	return
}
