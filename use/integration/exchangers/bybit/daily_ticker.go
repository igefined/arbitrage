package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	response "github.com/igilgyrg/arbitrage/use/integration/exchangers/bybit/response"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?symbol=%s&category=spot", "/v5/market/tickers", symbol)
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("bybit daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("bybit daily ticker response status: %s", resp.Status)

		return
	}

	responseBody := response.Response{}
	responseBody.Result = &response.TickerResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		err = fmt.Errorf("bybit daily ticker decoder: %v", err)

		return
	}

	if responseBody.Code != 0 {
		switch responseBody.Code {
		case 10001:
			err = fmt.Errorf("bybit daily ticker: %w", exchangers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("bybit daily ticker error response: %s", responseBody.Message)
		}

		return
	}

	if responseBody.Result == nil {
		err = fmt.Errorf("bybit daily ticker: nil result")

		return
	}

	tickerResponse, ok := responseBody.Result.(*response.TickerResponse)
	if !ok {
		err = fmt.Errorf("bybit daily ticker decoder: cannot json decode result")

		return
	}

	if tickerResponse.Category != "spot" {
		err = fmt.Errorf("bybit daily ticker: not spot account")

		return
	}

	if len(tickerResponse.List) == 0 {
		err = fmt.Errorf("bybit daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	ticker = tickerResponse.List[0].ToDomain()
	if ticker.Price <= 0 {
		err = fmt.Errorf("bybit ask price is zero for crypto %s", ticker.Symbol)

		return
	}

	return
}
