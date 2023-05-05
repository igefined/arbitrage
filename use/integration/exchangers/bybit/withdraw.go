package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/bybit/response"
	"github.com/igilgyrg/arbitrage/utils/usign"
)

func (c *client) WithdrawNetwork(ctx context.Context, symbol string) (networks []string) {
	queryParam := fmt.Sprintf("coin=%s", symbol)
	apiTimestamp := time.Now().UTC().UnixMilli()
	queryToHash := fmt.Sprintf("%d%s%d%s", apiTimestamp, c.cfg.ApiKey, recvWindow, queryParam)
	signatureHash := usign.SignRequest(queryToHash, c.cfg.SecretKey)

	query := fmt.Sprintf("%s?%s", "v5/asset/coin/query-info", queryParam)
	headers := map[string]string{
		"X-BAPI-API-KEY":     c.cfg.ApiKey,
		"X-BAPI-TIMESTAMP":   strconv.Itoa(int(apiTimestamp)),
		"X-BAPI-RECV-WINDOW": strconv.Itoa(recvWindow),
		"X-BAPI-SIGN":        signatureHash,
	}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		c.logger.Errorf("bybit all coins request: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		c.logger.Errorf("bybit all coins response status: %s", resp.Status)

		return
	}

	responseBody := response.Response{}
	responseBody.Result = &response.CoinInfoResp{}
	if err = json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		c.logger.Errorf("bybit query info decoder: %v", err)

		return
	}

	if responseBody.Code != 0 {
		switch responseBody.Code {
		case 10001:
			c.logger.Errorf("bybit query info: %w", exchangers.ErrSymbolNotFound)
		case 10002:
			c.logger.Errorf("bybit query info: %w", exchangers.ErrSymbolNotFound)
		default:
			c.logger.Errorf("bybit query info error response: %s", responseBody.Message)
		}

		return
	}

	if responseBody.Result == nil {
		c.logger.Errorf("bybit query info: nil result")

		return
	}

	coinInfoResponse, ok := responseBody.Result.(*response.CoinInfoResp)
	if !ok {
		c.logger.Errorf("bybit query info decoder: cannot json decode result")

		return
	}

	if len(coinInfoResponse.Rows) == 0 {
		c.logger.Errorf("bybit query info: %s", exchangers.ErrEmptyNetworks)

		return
	}

	if len(coinInfoResponse.Rows[0].Chains) == 0 {
		c.logger.Errorf("bybit query info: %s", exchangers.ErrEmptyNetworks)

		return
	}

	networks = make([]string, len(coinInfoResponse.Rows[0].Chains))
	for i, chain := range coinInfoResponse.Rows[0].Chains {
		if chain.ChainWithdraw == "1" {
			networks[i] = chain.Chain
		}
	}

	if err != nil {
		c.logger.Error(err)
	}

	return
}
