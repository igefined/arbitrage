package kucoin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/kucoin/response"
)

func (c *client) DepositNetwork(ctx context.Context, symbol string) (networks []string) {
	query := fmt.Sprintf("%s/%s", "api/v2/currencies", symbol)
	headers := map[string]string{}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		c.logger.Errorf("kucoin daily ticker request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		c.logger.Errorf("kucoin daily ticker response status: %s", resp.Status)

		return
	}

	responseBody := response.Response{}
	responseBody.Data = &response.CurrencyDetail{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		c.logger.Errorf("kucoin daily ticker decoder: %v", err)

		return
	}

	if responseBody.Data == nil {
		c.logger.Error("kucoin daily ticker: nil result")

		return
	}

	currency, ok := responseBody.Data.(*response.CurrencyDetail)
	if !ok {
		c.logger.Error("kucoin daily ticker decoder: cannot json decode result")

		return
	}

	if len(currency.Chains) == 0 {
		c.logger.Errorf("kucoin daily ticker: %v", exchangers.ErrSymbolNotFound)

		return
	}

	networks = make([]string, 0, len(currency.Chains))
	for _, chain := range currency.Chains {
		if chain.IsDepositEnabled {
			networks = append(networks, chain.Name)
		}
	}

	return
}
