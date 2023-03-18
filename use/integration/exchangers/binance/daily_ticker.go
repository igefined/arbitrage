package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?symbol=%s", "api/v3/ticker/24hr", symbol)
	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, nil)
	if err != nil {
		err = fmt.Errorf("binance daily ticker request: %v", err)

		return
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		errResp := &ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			err = fmt.Errorf("binance daily ticker decoder: %v", err)

			return
		}

		switch errResp.Code {
		case -1121:
			err = fmt.Errorf("binance daily ticker: %w", exchangers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("binance daily ticker error response: %v", errResp)
		}

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("binance daily ticker response status: %s", resp.Status)

		return
	}

	response := &DailyTickerResponse{}
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		err = fmt.Errorf("binance daily ticker decoder: %v", err)

		return
	}

	ticker = response.ToResponse()

	return
}
