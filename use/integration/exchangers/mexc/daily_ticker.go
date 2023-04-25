package mexc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/mexc/response"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?symbol=%s", "api/v3/ticker/bookTicker", symbol)
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("mexc daily ticker request: %v", err)

		return
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		errResp := &response.ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			err = fmt.Errorf("mexc daily ticker decoder: %v", err)

			return
		}

		switch errResp.Code {
		case -1121:
			err = fmt.Errorf("mexc daily ticker: %w", exchangers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("mexc daily ticker error response: %v", errResp)
		}

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("mexc daily ticker response status: %s", resp.Status)

		return
	}

	response := &response.DailyTickerResponse{}
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		err = fmt.Errorf("mexc daily ticker decoder: %v", err)

		return
	}

	ticker = response.ToResponse()
	if ticker.Price == 0 {
		err = fmt.Errorf("mexc ask price is zero for crypto %s", ticker.Symbol)

		return
	}

	return
}
