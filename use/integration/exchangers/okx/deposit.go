package okx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/okx/response"
	"github.com/igilgyrg/arbitrage/utils/usign"
)

func (c *client) DepositNetwork(ctx context.Context, symbol string) (networks []string) {
	query := fmt.Sprintf("%s?ccy=%s", "/api/v5/asset/deposit-address", symbol)
	now := time.Now().UTC().Add(time.Second * 10).Format("2006-01-02T15:04:05.000Z")

	queryToSign := now + http.MethodGet + query

	headers := map[string]string{
		"OK-ACCESS-KEY":        c.cfg.ApiKey,
		"OK-ACCESS-TIMESTAMP":  now,
		"OK-ACCESS-PASSPHRASE": c.cfg.Passphrase,
		"OK-ACCESS-SIGN":       usign.SignRequestBase64(queryToSign, c.cfg.SecretKey),
	}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		c.logger.Errorf("okx deposit addresses do request err: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		c.logger.Errorf("okx deposit addresses response error: Code:%d Status:%s", resp.StatusCode, resp.Status)

		return
	}

	responseBody := &response.Response{}
	responseBody.Data = &response.DepositAddresses{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		c.logger.Errorf("okx daily ticker decoder: %s", err)

		return
	}

	if responseBody.Code == "51001" {
		c.logger.Errorf("okx daily ticker: %s", exchangers.ErrSymbolNotFound)

		return
	}

	responseChains, ok := (responseBody.Data).(*response.DepositAddresses)
	if !ok {
		c.logger.Errorf("okx cannont cast to deposit chain addresses")

		return
	}

	if responseChains == nil {
		c.logger.Errorf("okx daily ticker: %s", exchangers.ErrSymbolNotFound)

		return
	}

	if len(*responseChains) == 0 {
		c.logger.Error("okx zero chains supported")

		return
	}

	networks = make([]string, 0, len(*responseChains))
	for _, ch := range *responseChains {
		networks = append(networks, ch.Chain)
	}

	return
}
