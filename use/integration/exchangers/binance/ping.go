package binance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/igilgyrg/arbitrage/use/integration/exchangers"
)

func (c *client) Ping(ctx context.Context) (err error) {
	resp, err := exchangers.DoRequest(ctx, c.httpClient, http.MethodGet, c.hosts, "/api/v3/ping", nil)
	if err != nil {
		err = fmt.Errorf("binance ping: %v", err)

		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("binance ping response status: %s", resp.Status)

		return
	}

	return
}
