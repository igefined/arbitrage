package huobi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	response2 "github.com/igilgyrg/arbitrage/use/integration/exchangers/huobi/response"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?symbol=%s", "market/detail/merged", strings.ToLower(symbol))
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("huobi daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("huobi daily ticker response status: %s", resp.Status)

		return
	}

	response := response2.Response{}
	response.Result = &response2.DailyTickerResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		err = fmt.Errorf("huobi daily ticker decoder: %v", err)

		return
	}

	if response.ErrCode != "" {
		switch response.ErrCode {
		case "invalid-parameter":
			err = fmt.Errorf("huobi daily ticker: %w", exchangers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("huobi daily ticker error response: %s", response.ErrMessage)
		}

		return
	}

	if response.Result == nil {
		err = fmt.Errorf("huobi daily ticker: nil result")

		return
	}

	tickerResponse, ok := response.Result.(*response2.DailyTickerResponse)
	if !ok {
		err = fmt.Errorf("huobi daily ticker decoder: cannot json decode result")

		return
	}

	if len(tickerResponse.Ask) == 0 {
		err = fmt.Errorf("huobi daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	ticker = tickerResponse.ToResponse(symbol)
	if ticker.Price <= 0 {
		err = fmt.Errorf("huobi ask price is zero for crypto %s", ticker.Symbol)

		return
	}

	return
}
