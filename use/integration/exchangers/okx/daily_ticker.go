package okx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/igilgyrg/arbitrage/use/domain"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/okx/response"
	"github.com/igilgyrg/arbitrage/utils/usymbol"
)

func (c *client) DailyTicker(ctx context.Context, symbol string) (ticker *domain.DailyTicker, err error) {
	query := fmt.Sprintf("%s?instId=%s", "api/v5/market/ticker", usymbol.SeparateSymbol(strings.ToUpper(symbol), "-"))
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("okx daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("okx daily ticker response status: %s", resp.Status)

		return
	}

	responseBody := response.DailyTicker{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		err = fmt.Errorf("okx daily ticker decoder: %v", err)

		return
	}

	if responseBody.Code == "51001" {
		err = fmt.Errorf("okx daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	if len(responseBody.Data) == 0 {
		err = fmt.Errorf("okx daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	responseTicker := (responseBody.Data)[0]

	if responseTicker.AskPx == "0" {
		err = fmt.Errorf("okx daily ticker: %w", exchangers.ErrSymbolNotFound)

		return
	}

	ticker = responseTicker.ToDomain()
	if ticker.Price <= 0 {
		err = fmt.Errorf("okx ask price is zero for crypto %s", ticker.Symbol)

		return
	}

	return
}
