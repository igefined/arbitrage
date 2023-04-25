package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
	"github.com/igilgyrg/arbitrage/use/integration/exchangers/binance/response"
	"github.com/igilgyrg/arbitrage/utils/usign"
)

func (c *client) AllCoinsInfoProcessed(ctx context.Context) (err error) {
	queryParam := fmt.Sprintf("recvWindow=%d&timestamp=%d", recvWindow, time.Now().UTC().UnixMilli())
	query := fmt.Sprintf("%s?%s&signature=%s", "sapi/v1/capital/config/getall", queryParam, usign.SignRequest(queryParam, c.cfg.SecretKey))
	headers := map[string]string{
		"X-MBX-APIKEY": c.cfg.ApiKey,
	}

	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, query, headers, nil)
	if err != nil {
		err = fmt.Errorf("binance all coins request: %v", err)

		return
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		errResp := &response.ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			err = fmt.Errorf("binance all coins decoder: %v", err)

			return
		}

		switch errResp.Code {
		case -1121:
			err = fmt.Errorf("binance all coins: %w", exchangers.ErrSymbolNotFound)
		default:
			err = fmt.Errorf("binance all coins error response: %v", errResp)
		}

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("binance all coins response status: %s", resp.Status)

		return
	}

	var allCoinsInformation []response.CoinInformation
	if err = json.NewDecoder(resp.Body).Decode(&allCoinsInformation); err != nil {
		err = fmt.Errorf("binance all coins decoder: %v", err)

		return
	}

	c.allCoinsInfo = make(map[string]response.CoinInformation, len(allCoinsInformation))
	for _, info := range allCoinsInformation {
		c.allCoinsInfo[info.Coin] = info
	}

	return
}
