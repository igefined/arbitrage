package kucoin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/kucoin/response"
	"github.com/igilgyrg/arbitrage/utils/usymbol"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?symbol=%s", "api/v1/market/stats", usymbol.SeparateSymbol(strings.ToUpper(symbol), "-"))
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("kucoin daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("kucoin daily ticker response status: %s", resp.Status)

		return
	}

	responseBody := response.Response{}
	responseBody.Data = &response.DailyTicker{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		err = fmt.Errorf("kucoin daily ticker decoder: %v", err)

		return
	}

	if responseBody.Data == nil {
		err = fmt.Errorf("kucoin daily ticker: nil result")

		return
	}

	tickerResponse, ok := responseBody.Data.(*response.DailyTicker)
	if !ok {
		err = fmt.Errorf("kucoin daily ticker decoder: cannot json decode result")

		return
	}

	if len(tickerResponse.Sell) == 0 {
		err = fmt.Errorf("kucoin daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	ticker = tickerResponse.ToDomain()
	if ticker.Price <= 0 {
		err = fmt.Errorf("kucoin ask price is zero for crypto %s", ticker.Symbol)

		return
	}

	return
}
