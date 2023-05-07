package mexc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/mexc/response"
	"github.com/igilgyrg/arbitrage/utils/usign"
)

func (c *client) WithdrawNetwork(ctx context.Context, symbol string) (networks []string) {
	now := time.Now().UnixMilli()
	queryToSign := fmt.Sprintf("coin=%s&timestamp=%d&recvWindow=30000", symbol, now)
	query := fmt.Sprintf("%s?%s&signature=%s", "/api/v3/capital/withdraw/address", queryToSign, usign.SignRequest(queryToSign, c.cfg.SecretKey))

	headers := map[string]string{
		"X-MEXC-APIKEY": c.cfg.ApiKey,
	}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		c.logger.Errorf("mexc daily ticker request: %v", err)

		return
	}

	if resp.StatusCode >= 400 {
		errResp := &response.ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			c.logger.Errorf("mexc deposit addresses decoder: %v", err)

			return
		}

		switch errResp.Code {
		case -1121:
			c.logger.Errorf("mexc deposit addresses: %v", exchangers.ErrSymbolNotFound)
		default:
			c.logger.Errorf("mexc deposit addresses error response: %v", errResp)
		}

		return
	}

	if resp.StatusCode != 200 {
		c.logger.Errorf("mexc deposit addresses response status: %s", resp.Status)

		return
	}

	var responseBody response.WithdrawAddresses
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		c.logger.Errorf("mexc deposit addresses decoder: %v", err)

		return
	}

	networks = make([]string, 0, len(responseBody.Data))
	for _, a := range responseBody.Data {
		networks = append(networks, a.Network)
	}

	return
}
